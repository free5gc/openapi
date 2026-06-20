package EventExposure

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/free5gc/openapi"
	"github.com/stretchr/testify/require"
)

func TestDeleteSubscriptionStatusHandling(t *testing.T) {
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
			status:  http.StatusNoContent,
			kind:    successResponse,
			headers: http.Header{},
		},
	}
	for _, status := range []int{400, 401, 403, 404, 411, 413, 415, 429, 500, 502, 503} {
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
				"Location":              []string{"https://redirect.example.com/nupf-ee/v1/ee-subscriptions/sub-1"},
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
			request := &DeleteSubscriptionRequest{}
			request.SetSubscriptionId("sub-1")

			response, err := client.IndividualSubscriptionDocumentApi.DeleteSubscription(context.Background(), request)

			require.Equal(t, http.MethodDelete, captured.method)
			require.Equal(t, "/nupf-ee/v1/ee-subscriptions/sub-1", captured.path)

			if test.kind == successResponse {
				require.NoError(t, err)
				require.NotNil(t, response)
				return
			}

			require.Nil(t, response)
			require.Error(t, err)
			var apiError openapi.GenericOpenAPIError
			require.ErrorAs(t, err, &apiError)
			require.Equal(t, test.status, apiError.ErrorStatus)

			switch test.kind {
			case problemResponse:
				model, ok := apiError.Model().(DeleteSubscriptionError)
				require.True(t, ok)
				require.Equal(t, int32(test.status), model.ProblemDetails.Status)
				require.Equal(t, "TEST_PROBLEM", model.ProblemDetails.Cause)
			case redirectResponse:
				model, ok := apiError.Model().(DeleteSubscriptionError)
				require.True(t, ok)
				require.Equal(t, "TEMPORARY_REDIRECTION", model.RedirectResponse.Cause)
				require.Equal(t, "https://redirect.example.com/nupf-ee/v1/ee-subscriptions/sub-1", model.Location)
				require.Equal(t, "target-nf-id", model.Var3gppSbiTargetNfId)
				require.Equal(t, 1, captured.requestCount)
				require.Equal(t, 1, captured.originCount)
				require.Zero(t, captured.redirectCount)
				require.Zero(t, captured.requestBodyCount)
			case defaultResponse:
				require.Nil(t, apiError.Model())
			}
		})
	}
}
