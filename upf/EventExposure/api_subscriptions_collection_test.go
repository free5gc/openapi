package EventExposure

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

type testRoundTripper func(*http.Request) (*http.Response, error)

func (f testRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

type capturedRequest struct {
	method           string
	path             string
	body             []byte
	requestCount     int
	originCount      int
	redirectCount    int
	requestBodyCount int
}

func newTestAPIClient(status int, body string, headers http.Header, policy openapi.RedirectPolicy) (*APIClient, *capturedRequest) {
	captured := &capturedRequest{}
	httpClient := &http.Client{
		Transport: testRoundTripper(func(request *http.Request) (*http.Response, error) {
			captured.requestCount++
			switch request.URL.Host {
			case "upf.example.com":
				captured.originCount++
			case "redirect.example.com":
				captured.redirectCount++
			}
			captured.method = request.Method
			captured.path = request.URL.Path
			if request.Body != nil {
				captured.requestBodyCount++
				requestBody, err := io.ReadAll(request.Body)
				if err != nil {
					return nil, err
				}
				captured.body = requestBody
			}

			responseHeaders := make(http.Header)
			for key, values := range headers {
				responseHeaders[key] = append([]string(nil), values...)
			}
			return &http.Response{
				StatusCode: status,
				Header:     responseHeaders,
				Body:       io.NopCloser(strings.NewReader(body)),
				Request:    request,
			}, nil
		}),
	}

	configuration := NewConfiguration()
	configuration.SetBasePath("https://upf.example.com")
	configuration.SetHTTPClient(httpClient)
	configuration.SetRedirectPolicy(policy)
	return NewAPIClient(configuration), captured
}

func TestCreateSubscriptionStatusHandling(t *testing.T) {
	type responseKind string
	const (
		successResponse  responseKind = "success"
		problemResponse  responseKind = "problem"
		redirectResponse responseKind = "redirect"
		defaultResponse  responseKind = "default"
	)
	type testCase struct {
		status  int
		kind    responseKind
		body    string
		headers http.Header
	}

	tests := []testCase{
		{
			status: http.StatusCreated,
			kind:   successResponse,
			body: `{
				"subscription":{
					"eventList":[{
						"type":"USER_DATA_USAGE_MEASURES",
						"measurementTypes":["VOLUME_MEASUREMENT"],
						"granularityOfMeasurement":"PER_SESSION"
					}],
					"eventNotifyUri":"https://nwdaf.example.com/upf-events",
					"notifyCorrelationId":"nsmf-notif-id",
					"eventReportingMode":{"trigger":"PERIODIC","repPeriod":10},
					"nfId":"11111111-2222-3333-4444-555555555555",
					"ueIpAddress":{"ipv4Addr":"192.0.2.1"}
				},
				"subscriptionId":"sub-1",
				"supportedFeatures":"1"
			}`,
			headers: http.Header{
				"Content-Type": []string{"application/json"},
				"Location":     []string{"https://upf.example.com/nupf-ee/v1/ee-subscriptions/sub-1"},
			},
		},
	}

	for _, status := range []int{400, 401, 403, 404, 411, 413, 415, 429, 500, 501, 502, 503} {
		tests = append(tests, testCase{
			status: status,
			kind:   problemResponse,
			body:   fmt.Sprintf(`{"status":%d,"cause":"TEST_PROBLEM"}`, status),
			headers: http.Header{
				"Content-Type": []string{"application/problem+json"},
			},
		})
	}
	for _, status := range []int{307, 308} {
		tests = append(tests, testCase{
			status: status,
			kind:   redirectResponse,
			body:   `{"cause":"TEMPORARY_REDIRECTION"}`,
			headers: http.Header{
				"Content-Type":          []string{"application/json"},
				"Location":              []string{"https://redirect.example.com/nupf-ee/v1/ee-subscriptions"},
				"3gpp-Sbi-Target-Nf-Id": []string{"target-nf-id"},
			},
		})
	}
	tests = append(tests, testCase{
		status:  http.StatusTeapot,
		kind:    defaultResponse,
		body:    "unexpected response",
		headers: http.Header{"Content-Type": []string{"text/plain"}},
	})

	for _, test := range tests {
		t.Run(fmt.Sprintf("status_%d", test.status), func(t *testing.T) {
			var policy openapi.RedirectPolicy
			if test.kind == redirectResponse {
				policy = openapi.RejectRedirects
			}
			client, captured := newTestAPIClient(test.status, test.body, test.headers, policy)
			request := &CreateSubscriptionRequest{}
			request.SetUpfCreateEventSubscription(testCreateEventSubscription())

			response, err := client.SubscriptionsCollectionApi.CreateSubscription(context.Background(), request)

			require.Equal(t, http.MethodPost, captured.method)
			require.Equal(t, "/nupf-ee/v1/ee-subscriptions", captured.path)
			require.JSONEq(t, `{
				"subscription":{
					"eventList":[{
						"type":"USER_DATA_USAGE_MEASURES",
						"measurementTypes":["VOLUME_MEASUREMENT"],
						"granularityOfMeasurement":"PER_SESSION"
					}],
					"eventNotifyUri":"https://nwdaf.example.com/upf-events",
					"notifyCorrelationId":"nsmf-notif-id",
					"eventReportingMode":{"trigger":"PERIODIC","repPeriod":10},
					"nfId":"11111111-2222-3333-4444-555555555555",
					"ueIpAddress":{"ipv4Addr":"192.0.2.1"}
				}
			}`, string(captured.body))

			if test.kind == successResponse {
				require.NoError(t, err)
				require.Equal(t, "https://upf.example.com/nupf-ee/v1/ee-subscriptions/sub-1", response.Location)
				created := response.UpfCreatedEventSubscription
				require.Equal(t, "sub-1", created.SubscriptionId)
				require.Equal(t, "1", created.SupportedFeatures)
				require.Equal(t, "https://nwdaf.example.com/upf-events", created.Subscription.EventNotifyUri)
				require.Equal(t, "nsmf-notif-id", created.Subscription.NotifyCorrelationId)
				require.Equal(t, models.UpfEventTrigger_PERIODIC, created.Subscription.EventReportingMode.Trigger)
				require.Equal(t, int32(10), created.Subscription.EventReportingMode.RepPeriod)
				require.Equal(t, "11111111-2222-3333-4444-555555555555", created.Subscription.NfId)
				require.Equal(t, &models.IpAddr{Ipv4Addr: "192.0.2.1"}, created.Subscription.UeIpAddress)
				require.Len(t, created.Subscription.EventList, 1)
				require.Equal(t, models.UpfEventType_USER_DATA_USAGE_MEASURES, created.Subscription.EventList[0].Type)
				require.Equal(t, []models.UpfMeasurementType{models.UpfMeasurementType_VOLUME_MEASUREMENT}, created.Subscription.EventList[0].MeasurementTypes)
				require.Equal(t, models.UpfGranularityOfMeasurement_PER_SESSION, created.Subscription.EventList[0].GranularityOfMeasurement)
				return
			}

			require.Nil(t, response)
			require.Error(t, err)
			var apiError openapi.GenericOpenAPIError
			require.ErrorAs(t, err, &apiError)
			require.Equal(t, test.status, apiError.ErrorStatus)

			switch test.kind {
			case problemResponse:
				model, ok := apiError.Model().(CreateSubscriptionError)
				require.True(t, ok)
				require.Equal(t, int32(test.status), model.ProblemDetails.Status)
				require.Equal(t, "TEST_PROBLEM", model.ProblemDetails.Cause)
			case redirectResponse:
				model, ok := apiError.Model().(CreateSubscriptionError)
				require.True(t, ok)
				require.Equal(t, "TEMPORARY_REDIRECTION", model.RedirectResponse.Cause)
				require.Equal(t, "https://redirect.example.com/nupf-ee/v1/ee-subscriptions", model.Location)
				require.Equal(t, "target-nf-id", model.Var3gppSbiTargetNfId)
				require.Equal(t, 1, captured.requestCount)
				require.Equal(t, 1, captured.originCount)
				require.Zero(t, captured.redirectCount)
				require.Equal(t, 1, captured.requestBodyCount)
			case defaultResponse:
				require.Nil(t, apiError.Model())
			}
		})
	}
}

func TestCreateSubscriptionFollowsRedirectWithReplayableBody(t *testing.T) {
	type observedRequest struct {
		method string
		host   string
		path   string
		body   []byte
	}

	for _, redirectStatus := range []int{http.StatusTemporaryRedirect, http.StatusPermanentRedirect} {
		t.Run(fmt.Sprintf("status_%d", redirectStatus), func(t *testing.T) {
			var observed []observedRequest
			httpClient := &http.Client{
				Transport: testRoundTripper(func(request *http.Request) (*http.Response, error) {
					requestBody, err := io.ReadAll(request.Body)
					if err != nil {
						return nil, err
					}
					observed = append(observed, observedRequest{
						method: request.Method,
						host:   request.URL.Host,
						path:   request.URL.Path,
						body:   requestBody,
					})

					if len(observed) == 1 {
						return &http.Response{
							StatusCode: redirectStatus,
							Header: http.Header{
								"Content-Type": []string{"application/json"},
								"Location":     []string{"https://redirect.example.com/nupf-ee/v1/ee-subscriptions"},
							},
							Body:    io.NopCloser(strings.NewReader(`{"cause":"TEMPORARY_REDIRECTION"}`)),
							Request: request,
						}, nil
					}

					return &http.Response{
						StatusCode: http.StatusCreated,
						Header: http.Header{
							"Content-Type": []string{"application/json"},
							"Location":     []string{"https://redirect.example.com/nupf-ee/v1/ee-subscriptions/sub-1"},
						},
						Body: io.NopCloser(strings.NewReader(`{
							"subscription":{
								"eventList":[{"type":"USER_DATA_USAGE_MEASURES"}],
								"eventNotifyUri":"https://nwdaf.example.com/upf-events",
								"notifyCorrelationId":"nsmf-notif-id",
								"eventReportingMode":{"trigger":"PERIODIC"},
								"nfId":"11111111-2222-3333-4444-555555555555"
							},
							"subscriptionId":"sub-1"
						}`)),
						Request: request,
					}, nil
				}),
			}

			configuration := NewConfiguration()
			configuration.SetBasePath("https://upf.example.com")
			configuration.SetHTTPClient(httpClient)
			client := NewAPIClient(configuration)
			request := &CreateSubscriptionRequest{}
			request.SetUpfCreateEventSubscription(testCreateEventSubscription())

			response, err := client.SubscriptionsCollectionApi.CreateSubscription(context.Background(), request)

			require.NoError(t, err)
			require.Equal(t, "sub-1", response.UpfCreatedEventSubscription.SubscriptionId)
			require.Equal(t, "https://redirect.example.com/nupf-ee/v1/ee-subscriptions/sub-1", response.Location)
			require.Len(t, observed, 2)
			require.Equal(t, http.MethodPost, observed[0].method)
			require.Equal(t, http.MethodPost, observed[1].method)
			require.Equal(t, "upf.example.com", observed[0].host)
			require.Equal(t, "redirect.example.com", observed[1].host)
			require.Equal(t, "/nupf-ee/v1/ee-subscriptions", observed[0].path)
			require.Equal(t, observed[0].path, observed[1].path)
			require.JSONEq(t, string(observed[0].body), string(observed[1].body))
		})
	}
}

func testCreateEventSubscription() models.UpfCreateEventSubscription {
	return models.UpfCreateEventSubscription{
		Subscription: models.UpfEventSubscription{
			EventList: []models.UpfEvent{
				{
					Type:                     models.UpfEventType_USER_DATA_USAGE_MEASURES,
					MeasurementTypes:         []models.UpfMeasurementType{models.UpfMeasurementType_VOLUME_MEASUREMENT},
					GranularityOfMeasurement: models.UpfGranularityOfMeasurement_PER_SESSION,
				},
			},
			EventNotifyUri:      "https://nwdaf.example.com/upf-events",
			NotifyCorrelationId: "nsmf-notif-id",
			EventReportingMode: models.UpfEventMode{
				Trigger:   models.UpfEventTrigger_PERIODIC,
				RepPeriod: 10,
			},
			NfId:        "11111111-2222-3333-4444-555555555555",
			UeIpAddress: &models.IpAddr{Ipv4Addr: "192.0.2.1"},
		},
	}
}
