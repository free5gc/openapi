package openapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type clientTestRoundTripper func(*http.Request) (*http.Response, error)

func (f clientTestRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

type clientTestConfiguration struct {
	httpClient *http.Client
	metrics    RequestMetricsHook
}

func (c *clientTestConfiguration) BasePath() string                 { return "" }
func (c *clientTestConfiguration) Host() string                     { return "" }
func (c *clientTestConfiguration) UserAgent() string                { return "" }
func (c *clientTestConfiguration) DefaultHeader() map[string]string { return nil }
func (c *clientTestConfiguration) HTTPClient() *http.Client         { return c.httpClient }
func (c *clientTestConfiguration) Metrics() RequestMetricsHook      { return c.metrics }

type redirectClientTestConfiguration struct {
	*clientTestConfiguration
	policy RedirectPolicy
}

func (c *redirectClientTestConfiguration) RedirectPolicy() RedirectPolicy {
	return c.policy
}

func TestParameterToString(t *testing.T) {
	var testCases = []struct {
		name   string
		obj    interface{}
		format string
		para   string
	}{
		{
			name:   "strings csv",
			obj:    []string{"red", "blue", "yellow"},
			format: "csv",
			para:   "red,blue,yellow",
		},
		{
			name:   "strings pipes",
			obj:    []string{"red", "blue", "yellow"},
			format: "pipes",
			para:   "red|blue|yellow",
		},
		{
			name:   "strings ssv",
			obj:    []string{"red", "blue", "yellow"},
			format: "ssv",
			para:   "red blue yellow",
		},
		{
			name:   "strings json",
			obj:    []string{"red", "blue", "yellow"},
			format: "application/json",
			para:   "[\"red\",\"blue\",\"yellow\"]",
		},
		{
			name: "obj json",
			obj: struct {
				Name   string `json:"name"`
				Age    int    `json:"age"`
				Active bool   `json:"active"`
			}{
				Name:   "free5GC",
				Age:    3,
				Active: true,
			},
			format: "application/json",
			para:   "{\"name\":\"free5GC\",\"age\":3,\"active\":true}",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			para := ParameterToString(tc.obj, tc.format)
			require.Equal(t, tc.para, para)
		})
	}
}

// TestMultipartDeserialize_MalformedBody pins the fix for free5gc/free5gc#1026.
// Before the fix, a body whose Content-Type declared multipart/related but
// whose payload was not a valid MIME multipart stream would panic on
// part.Header.Get in MultipartDeserialize because mime/multipart.NextPart can
// return a nil *Part alongside a non-EOF error. The SMF then surfaced HTTP
// 500 on every malformed POST /nsmf-pdusession/v1/sm-contexts.
func TestMultipartDeserialize_MalformedBody(t *testing.T) {
	type dummy struct{}

	cases := []struct {
		name string
		body []byte
	}{
		{name: "plain text body", body: []byte("hello, not multipart")},
		{name: "empty body", body: []byte("")},
		{name: "json body", body: []byte(`{"foo":"bar"}`)},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var out dummy
			require.NotPanics(t, func() {
				err := MultipartDeserialize(tc.body, &out, "boundaryX")
				require.Error(t, err)
			})
		})
	}
}

func TestCallAPIPreservesFollowBehaviorWithoutPolicy(t *testing.T) {
	tests := []struct {
		name      string
		newConfig func(*http.Client) Configuration
	}{
		{
			name: "provider absent",
			newConfig: func(client *http.Client) Configuration {
				return &clientTestConfiguration{httpClient: client}
			},
		},
		{
			name: "nil policy",
			newConfig: func(client *http.Client) Configuration {
				return &redirectClientTestConfiguration{
					clientTestConfiguration: &clientTestConfiguration{httpClient: client},
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var requests atomic.Int32
			client := &http.Client{
				Transport: clientTestRoundTripper(func(request *http.Request) (*http.Response, error) {
					requests.Add(1)
					if request.URL.Host == "origin.example.com" {
						return clientTestResponse(request, http.StatusTemporaryRedirect, http.Header{
							"Location": []string{"https://target.example.com/resource"},
						}), nil
					}
					return clientTestResponse(request, http.StatusNoContent, nil), nil
				}),
			}
			request, err := http.NewRequest(http.MethodGet, "https://origin.example.com/resource", nil)
			require.NoError(t, err)

			response, err := CallAPI(test.newConfig(client), request)

			require.NoError(t, err)
			require.Equal(t, http.StatusNoContent, response.StatusCode)
			require.NoError(t, response.Body.Close())
			require.Equal(t, int32(2), requests.Load())
			require.Nil(t, client.CheckRedirect)
		})
	}
}

func TestSelectHTTPClientPreservesSelectedClient(t *testing.T) {
	t.Run("explicit client takes precedence", func(t *testing.T) {
		transport := clientTestRoundTripper(func(request *http.Request) (*http.Response, error) {
			return clientTestResponse(request, http.StatusNoContent, nil), nil
		})
		jar, err := cookiejar.New(nil)
		require.NoError(t, err)
		originalRedirectError := errors.New("original redirect policy")
		originalRedirectPolicy := func(_ *http.Request, _ []*http.Request) error {
			return originalRedirectError
		}
		original := &http.Client{
			Transport:     transport,
			CheckRedirect: originalRedirectPolicy,
			Jar:           jar,
			Timeout:       23 * time.Second,
		}
		configuration := &redirectClientTestConfiguration{
			clientTestConfiguration: &clientTestConfiguration{httpClient: original},
			policy:                  RejectRedirects,
		}
		request, err := http.NewRequest(http.MethodGet, "custom://origin.example.com/resource", nil)
		require.NoError(t, err)

		selected, err := selectHTTPClient(configuration, request)

		require.NoError(t, err)
		require.NotSame(t, original, selected)
		requireSameHTTPClientSettings(t, original, selected)
		require.ErrorIs(t, selected.CheckRedirect(nil, nil), http.ErrUseLastResponse)
		require.ErrorIs(t, original.CheckRedirect(nil, nil), originalRedirectError)
	})

	for _, test := range []struct {
		name     string
		scheme   string
		original *http.Client
	}{
		{name: "shared HTTPS client", scheme: "https", original: innerHTTP2Client},
		{name: "shared cleartext client", scheme: "http", original: innerHTTP2CleartextClient},
	} {
		t.Run(test.name, func(t *testing.T) {
			configuration := &redirectClientTestConfiguration{
				clientTestConfiguration: &clientTestConfiguration{},
				policy:                  RejectRedirects,
			}
			request, err := http.NewRequest(http.MethodGet, test.scheme+"://origin.example.com/resource", nil)
			require.NoError(t, err)

			selected, err := selectHTTPClient(configuration, request)

			require.NoError(t, err)
			require.NotSame(t, test.original, selected)
			requireSameHTTPClientSettings(t, test.original, selected)
			require.ErrorIs(t, selected.CheckRedirect(nil, nil), http.ErrUseLastResponse)
			require.Nil(t, test.original.CheckRedirect)
		})
	}
}

func TestCallAPIRedirectPolicyPreservesMetrics(t *testing.T) {
	var metricCalls atomic.Int32
	var metricStatus atomic.Int32
	var metricMethod string
	var metricPath string
	client := &http.Client{
		Transport: clientTestRoundTripper(func(request *http.Request) (*http.Response, error) {
			return clientTestResponse(request, http.StatusTemporaryRedirect, http.Header{
				"Location": []string{"https://target.example.com/resource"},
			}), nil
		}),
	}
	configuration := &redirectClientTestConfiguration{
		clientTestConfiguration: &clientTestConfiguration{
			httpClient: client,
			metrics: func(method string, path string, status int, _ float64) {
				metricCalls.Add(1)
				metricStatus.Store(int32(status))
				metricMethod = method
				metricPath = path
			},
		},
		policy: RejectRedirects,
	}
	request, err := http.NewRequest(http.MethodGet, "https://origin.example.com/nupf-ee/v1/resource", nil)
	require.NoError(t, err)

	response, err := CallAPI(configuration, request)

	require.NoError(t, err)
	require.Equal(t, http.StatusTemporaryRedirect, response.StatusCode)
	require.NoError(t, response.Body.Close())
	require.Equal(t, int32(1), metricCalls.Load())
	require.Equal(t, int32(http.StatusTemporaryRedirect), metricStatus.Load())
	require.Equal(t, http.MethodGet, metricMethod)
	require.Equal(t, "nupf-ee", metricPath)
	require.Nil(t, client.CheckRedirect)
}

func TestCallAPIRedirectPolicyConcurrentUse(t *testing.T) {
	client := &http.Client{
		Transport: clientTestRoundTripper(func(request *http.Request) (*http.Response, error) {
			if request.URL.Host == "origin.example.com" {
				return clientTestResponse(request, http.StatusTemporaryRedirect, http.Header{
					"Location": []string{"https://target.example.com/resource"},
				}), nil
			}
			return clientTestResponse(request, http.StatusNoContent, nil), nil
		}),
		Timeout: 5 * time.Second,
	}
	followConfiguration := &clientTestConfiguration{httpClient: client}
	rejectConfiguration := &redirectClientTestConfiguration{
		clientTestConfiguration: &clientTestConfiguration{httpClient: client},
		policy:                  RejectRedirects,
	}

	const iterations = 50
	errorsCh := make(chan error, iterations*2)
	var waitGroup sync.WaitGroup
	for range iterations {
		waitGroup.Add(2)
		go func() {
			defer waitGroup.Done()
			errorsCh <- callAPIExpectStatus(followConfiguration, http.StatusNoContent)
		}()
		go func() {
			defer waitGroup.Done()
			errorsCh <- callAPIExpectStatus(rejectConfiguration, http.StatusTemporaryRedirect)
		}()
	}
	waitGroup.Wait()
	close(errorsCh)

	for err := range errorsCh {
		require.NoError(t, err)
	}
	require.Nil(t, client.CheckRedirect)
}

func TestSelectHTTPClientRedirectPolicyConcurrentSharedUse(t *testing.T) {
	followConfiguration := &clientTestConfiguration{}
	rejectConfiguration := &redirectClientTestConfiguration{
		clientTestConfiguration: &clientTestConfiguration{},
		policy:                  RejectRedirects,
	}
	request, err := http.NewRequest(http.MethodGet, "https://origin.example.com/resource", nil)
	require.NoError(t, err)

	const iterations = 100
	errorsCh := make(chan error, iterations*2)
	var waitGroup sync.WaitGroup
	for range iterations {
		waitGroup.Add(2)
		go func() {
			defer waitGroup.Done()
			selected, selectErr := selectHTTPClient(followConfiguration, request)
			if selectErr != nil {
				errorsCh <- selectErr
				return
			}
			if selected != innerHTTP2Client || selected.CheckRedirect != nil {
				errorsCh <- errors.New("follow selection changed shared HTTPS client")
				return
			}
			errorsCh <- nil
		}()
		go func() {
			defer waitGroup.Done()
			selected, selectErr := selectHTTPClient(rejectConfiguration, request)
			if selectErr != nil {
				errorsCh <- selectErr
				return
			}
			if selected == innerHTTP2Client {
				errorsCh <- errors.New("reject selection returned shared HTTPS client")
				return
			}
			if !errors.Is(selected.CheckRedirect(nil, nil), http.ErrUseLastResponse) {
				errorsCh <- errors.New("reject selection did not install redirect policy")
				return
			}
			errorsCh <- nil
		}()
	}
	waitGroup.Wait()
	close(errorsCh)

	for err := range errorsCh {
		require.NoError(t, err)
	}
	require.Nil(t, innerHTTP2Client.CheckRedirect)
}

func callAPIExpectStatus(configuration Configuration, expectedStatus int) error {
	request, err := http.NewRequest(http.MethodGet, "https://origin.example.com/resource", nil)
	if err != nil {
		return err
	}
	response, err := CallAPI(configuration, request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != expectedStatus {
		return fmt.Errorf("status = %d, want %d", response.StatusCode, expectedStatus)
	}
	return nil
}

func clientTestResponse(request *http.Request, status int, header http.Header) *http.Response {
	if header == nil {
		header = make(http.Header)
	}
	return &http.Response{
		StatusCode: status,
		Header:     header,
		Body:       io.NopCloser(strings.NewReader("")),
		Request:    request,
	}
}

func requireSameHTTPClientSettings(t *testing.T, expected *http.Client, actual *http.Client) {
	t.Helper()
	require.Equal(t, reflect.ValueOf(expected.Transport).Pointer(), reflect.ValueOf(actual.Transport).Pointer())
	require.Equal(t, expected.Timeout, actual.Timeout)
	if expected.Jar == nil {
		require.Nil(t, actual.Jar)
	} else {
		require.Same(t, expected.Jar, actual.Jar)
	}
}
