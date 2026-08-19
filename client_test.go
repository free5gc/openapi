package openapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type clientTestRoundTripper func(*http.Request) (*http.Response, error)

func (f clientTestRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

type clientTestTransport struct {
	marker string
}

func (t *clientTestTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	return clientTestResponse(request, http.StatusNoContent, nil), nil
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

// Regression test: all HTTP clients must wrap their transport with otelhttp.NewTransport
// so that trace IDs are propagated in outgoing requests. If this breaks,
// downstream services will not receive the traceparent header and cross-service tracing fails.
func TestTransport_WrapsWithOtelTransport(t *testing.T) {
	t.Run("HTTP2CleartextClient.SetTransport", func(t *testing.T) {
		h2c := &HTTP2CleartextClient{client: &http.Client{}}
		h2c.SetTransport(1*time.Second, 2*time.Second, nil)

		_, ok := h2c.client.Transport.(*otelhttp.Transport)
		require.True(t, ok,
			"Transport must be *otelhttp.Transport to propagate trace IDs; "+
				"raw http2.Transport was set instead")
	})

	t.Run("NewHttp2Client", func(t *testing.T) {
		c := NewHttp2Client(false, 5*time.Second, 1*time.Second, 2*time.Second)

		_, ok := c.Transport.(*otelhttp.Transport)
		require.True(t, ok,
			"Transport must be *otelhttp.Transport to propagate trace IDs; "+
				"raw http2.Transport was set instead")
	})
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

func getStrPtr(s string) *string {
	return &s
}

func TestAddQueryParams(t *testing.T) {
	var testCases = []struct {
		name             string
		obj              any
		format           string
		queryParamString string
	}{
		{
			name:             "string multi",
			obj:              "red",
			format:           "multi",
			queryParamString: "color=red",
		},
		{
			name:             "string ptr multi",
			obj:              getStrPtr("red"),
			format:           "multi",
			queryParamString: "color=red",
		},
		{
			name:             "strings multi",
			obj:              []string{"red", "blue", "yellow"},
			format:           "multi",
			queryParamString: "color=red&color=blue&color=yellow",
		},
		{
			name:             "strings csv",
			obj:              []string{"red", "blue", "yellow"},
			format:           "csv",
			queryParamString: "color=" + url.QueryEscape("red,blue,yellow"),
		},
		{
			name:             "strings pipes",
			obj:              []string{"red", "blue", "yellow"},
			format:           "pipes",
			queryParamString: "color=" + url.QueryEscape("red|blue|yellow"),
		},
		{
			name:             "strings ssv",
			obj:              []string{"red", "blue", "yellow"},
			format:           "ssv",
			queryParamString: "color=" + url.QueryEscape("red blue yellow"),
		},
		{
			name:             "strings json",
			obj:              []string{"red", "blue", "yellow"},
			format:           "application/json",
			queryParamString: "color=" + url.QueryEscape(`["red","blue","yellow"]`),
		},
		{
			name: "obj json",
			obj: struct {
				Red    int `json:"red"`
				Blue   int `json:"blue"`
				Yellow int `json:"yellow"`
			}{
				Red:    255,
				Blue:   255,
				Yellow: 255,
			},
			format:           "application/json",
			queryParamString: "color=" + url.QueryEscape(`{"red":255,"blue":255,"yellow":255}`),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			queryParams := url.Values{}
			err := AddQueryParams(&queryParams, "color", tc.obj, tc.format)
			require.NoError(t, err)
			require.Equal(t, tc.queryParamString, queryParams.Encode())
		})
	}
}

func TestMultipartDeserialize_LargerThanBuffer(t *testing.T) {
	reqJsonBody := `{
		"n1MessageContainer":{
			"n1MessageClass":"5GMM",
			"n1MessageContent":{
				"contentId":"n1Msg"
			}
		},
		"registrationCtxtContainer":{
			"ueContext":{
				"supi":"imsi-2089300007487",
				"pei":"imeisv-1110000000000000",
				"seafData":{
					"ngKsi":{
						"tsc":"NATIVE",
						"ksi":0
					},
					"keyAmf":{
						"keyType":"",
						"keyVal":"d04d0dfb241eee4e7fc070866ee9be5724e74a362efade8b6d3e84defd03195c"
					},
					"nh":"f6cdd59292429570efc9558a548f244a8bb542e75b4b48acc59080755432fa34",
					"ncc":1
				}
			},
			"anType":"3GPP_ACCESS",
			"anN2ApId":1,
			"ranNodeId":{
				"plmnId":{
					"mcc":"208",
					"mnc":"93"
				},
				"gNbId":{
					"bitLength":22,
					"gNBValue":"000001"
				}
			},
			"initialAmfName":"amf1",
			"userLocation":{
				"nrLocation":{
					"tai":{
						"plmnId":{
							"mcc":"208",
							"mnc":"93"
						},
						"tac":"000001"
					},
					"ncgi":{
						"plmnId":{
							"mcc":"208",
							"mnc":"93"
						},
						"nrCellId":"000004001"
					},
					"ueLocationTimestamp":"2024-04-12T09:54:54.152026945Z"
				}
			},
			"rrcEstCause":"2",
			"anN2IPv4Addr":"172.16.21.2:39249",
			"allowedNssai":{
				"allowedSnssaiList":[
					{
						"allowedSnssai":{
							"sst":1,
							"sd":"111111"
						},
						"nsiInformationList":[
							{
								"nrfId":"http://10.100.200.10:8000/nnrf-nfm/v1/nf-instances",
								"nsiId":"2"
							}
						]
					}
				],
				"accessType":"3GPP_ACCESS"
			}
		}
	}`
	reqBody := "--D4/?XquckV=bx9{dC}EG[TUQEzEb.'(CG(k8,B:D++6w3M5qJ?P]B]-Xk" +
		"\r\n" +
		"Content-Type: application/json" +
		"\r\n\r\n" +
		reqJsonBody +
		"\r\n" +
		"--D4/?XquckV=bx9{dC}EG[TUQEzEb.'(CG(k8,B:D++6w3M5qJ?P]B]-Xk--" +
		"\r\n"
	fmt.Println(reqBody)
	notify := models.N1MessageNotifyRequestBody{}
	boundary := "D4/?XquckV=bx9{dC}EG[TUQEzEb.'(CG(k8,B:D++6w3M5qJ?P]B]-Xk"
	err := multipart.Unmarshal(boundary, []byte(reqBody), &notify)
	require.NoError(t, err)
	require.NotNil(t, notify.JsonData)
	require.NotNil(t, notify.JsonData.RegistrationCtxtContainer)
	require.NotNil(t, notify.JsonData.N1MessageContainer)
}

func TestCallAPIPreservesDefaultRedirectBehaviorWithoutPolicy(t *testing.T) {
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
							"Location": {"https://target.example.com/resource"},
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

func TestCallAPIPreservesExistingRedirectPolicyWithoutNewPolicy(t *testing.T) {
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
			var redirectCalls atomic.Int32
			var targetRequests atomic.Int32
			existingPolicy := func(_ *http.Request, _ []*http.Request) error {
				redirectCalls.Add(1)
				return http.ErrUseLastResponse
			}
			client := &http.Client{
				Transport: clientTestRoundTripper(func(request *http.Request) (*http.Response, error) {
					if request.URL.Host == "target.example.com" {
						targetRequests.Add(1)
						return clientTestResponse(request, http.StatusNoContent, nil), nil
					}
					return clientTestResponse(request, http.StatusTemporaryRedirect, http.Header{
						"Location": {"https://target.example.com/resource"},
					}), nil
				}),
				CheckRedirect: existingPolicy,
			}
			configuration := test.newConfig(client)
			request, err := http.NewRequest(http.MethodGet, "https://origin.example.com/resource", nil)
			require.NoError(t, err)

			selected, err := selectHTTPClient(configuration, request)
			require.NoError(t, err)
			require.Same(t, client, selected)

			response, err := CallAPI(configuration, request)
			require.NoError(t, err)
			require.Equal(t, http.StatusTemporaryRedirect, response.StatusCode)
			require.NoError(t, response.Body.Close())
			require.Equal(t, int32(1), redirectCalls.Load())
			require.Zero(t, targetRequests.Load())
			require.NotNil(t, client.CheckRedirect)
		})
	}
}

func TestSelectHTTPClientPreservesSelectedClient(t *testing.T) {
	t.Run("explicit client takes precedence", func(t *testing.T) {
		transport := &clientTestTransport{marker: "unchanged"}
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
		require.Equal(t, "unchanged", transport.marker)
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
			require.Nil(t, test.original.CheckRedirect)
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
	for _, redirectStatus := range []int{http.StatusTemporaryRedirect, http.StatusPermanentRedirect} {
		t.Run(fmt.Sprintf("status_%d", redirectStatus), func(t *testing.T) {
			var metricCalls atomic.Int32
			var metricStatus atomic.Int32
			var metricMethod string
			var metricService string
			client := &http.Client{
				Transport: clientTestRoundTripper(func(request *http.Request) (*http.Response, error) {
					return clientTestResponse(request, redirectStatus, http.Header{
						"Location": {"https://target.example.com/resource"},
					}), nil
				}),
			}
			configuration := &redirectClientTestConfiguration{
				clientTestConfiguration: &clientTestConfiguration{
					httpClient: client,
					metrics: func(method string, service string, status int, duration float64) {
						metricCalls.Add(1)
						metricStatus.Store(int32(status))
						metricMethod = method
						metricService = service
						require.GreaterOrEqual(t, duration, float64(0))
					},
				},
				policy: RejectRedirects,
			}
			request, err := http.NewRequest(http.MethodPost, "https://origin.example.com/nupf-ee/v1/resource", nil)
			require.NoError(t, err)

			response, err := CallAPI(configuration, request)

			require.NoError(t, err)
			require.Equal(t, redirectStatus, response.StatusCode)
			require.NoError(t, response.Body.Close())
			require.Equal(t, int32(1), metricCalls.Load())
			require.Equal(t, int32(redirectStatus), metricStatus.Load())
			require.Equal(t, http.MethodPost, metricMethod)
			require.Equal(t, "nupf-ee", metricService)
			require.Nil(t, client.CheckRedirect)
		})
	}
}

func TestCallAPIRedirectPolicyConcurrentUse(t *testing.T) {
	transport := clientTestRoundTripper(func(request *http.Request) (*http.Response, error) {
		if request.URL.Host == "origin.example.com" {
			return clientTestResponse(request, http.StatusTemporaryRedirect, http.Header{
				"Location": {"https://target.example.com/resource"},
			}), nil
		}
		return clientTestResponse(request, http.StatusNoContent, nil), nil
	})
	followClient := &http.Client{
		Transport: transport,
		Timeout:   5 * time.Second,
	}
	var existingPolicyCalls atomic.Int32
	existingPolicyClient := &http.Client{
		Transport: transport,
		Timeout:   5 * time.Second,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			existingPolicyCalls.Add(1)
			return http.ErrUseLastResponse
		},
	}
	followConfiguration := &clientTestConfiguration{httpClient: followClient}
	existingPolicyConfiguration := &clientTestConfiguration{httpClient: existingPolicyClient}
	rejectConfiguration := &redirectClientTestConfiguration{
		clientTestConfiguration: &clientTestConfiguration{httpClient: existingPolicyClient},
		policy:                  RejectRedirects,
	}

	const iterations = 50
	errorsCh := make(chan error, iterations*3)
	var waitGroup sync.WaitGroup
	for range iterations {
		waitGroup.Add(3)
		go func() {
			defer waitGroup.Done()
			errorsCh <- callAPIExpectStatus(followConfiguration, http.StatusNoContent)
		}()
		go func() {
			defer waitGroup.Done()
			errorsCh <- callAPIExpectStatus(existingPolicyConfiguration, http.StatusTemporaryRedirect)
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
	require.Nil(t, followClient.CheckRedirect)
	require.NotNil(t, existingPolicyClient.CheckRedirect)
	require.Equal(t, int32(iterations), existingPolicyCalls.Load())
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
	require.Same(t, expected.Transport, actual.Transport)
	require.Equal(t, expected.Timeout, actual.Timeout)
	if expected.Jar == nil {
		require.Nil(t, actual.Jar)
	} else {
		require.Same(t, expected.Jar, actual.Jar)
	}
}
