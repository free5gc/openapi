package EvtExpos

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

const redirectTargetNFID = "target-nf-id"

type redirectTestCapture struct {
	originRequests atomic.Int32
	targetRequests atomic.Int32
	originBodies   atomic.Int32

	mu           sync.Mutex
	originMethod string
	originPath   string
	originBody   string
}

type redirectTestSnapshot struct {
	originRequests int32
	targetRequests int32
	originBodies   int32
	originMethod   string
	originPath     string
	originBody     string
}

func (c *redirectTestCapture) recordOrigin(request *http.Request, body []byte) {
	c.originRequests.Add(1)
	if len(body) > 0 {
		c.originBodies.Add(1)
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.originMethod = request.Method
	c.originPath = request.URL.Path
	c.originBody = string(body)
}

func (c *redirectTestCapture) snapshot() redirectTestSnapshot {
	c.mu.Lock()
	defer c.mu.Unlock()
	return redirectTestSnapshot{
		originRequests: c.originRequests.Load(),
		targetRequests: c.targetRequests.Load(),
		originBodies:   c.originBodies.Load(),
		originMethod:   c.originMethod,
		originPath:     c.originPath,
		originBody:     c.originBody,
	}
}

func newRedirectTestAPIClient(
	t *testing.T,
	status int,
	locationPath string,
) (*APIClient, *redirectTestCapture, string) {
	t.Helper()

	capture := &redirectTestCapture{}
	targetServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		capture.targetRequests.Add(1)
		w.WriteHeader(http.StatusNoContent)
	}))
	t.Cleanup(targetServer.Close)

	location := targetServer.URL + locationPath
	originServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		capture.recordOrigin(request, body)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Location", location)
		w.Header().Set("3gpp-Sbi-Target-Nf-Id", redirectTargetNFID)
		w.WriteHeader(status)
		_, _ = io.WriteString(w, `{"cause":"TEMPORARY_REDIRECTION"}`)
	}))
	t.Cleanup(originServer.Close)

	httpClient := &http.Client{}
	require.Nil(t, httpClient.Transport, "redirect opt-in must not require a custom transport")
	configuration := NewConfiguration()
	configuration.SetBasePath(originServer.URL)
	configuration.SetHTTPClient(httpClient)
	require.Nil(t, configuration.RedirectPolicy())
	configuration.SetRedirectPolicy(openapi.RejectRedirects)
	require.ErrorIs(t, configuration.RedirectPolicy()(nil, nil), http.ErrUseLastResponse)

	return NewAPIClient(configuration), capture, location
}

func TestCreateSubscriptionRejectsRedirectWithoutReplay(t *testing.T) {
	for _, status := range []int{http.StatusTemporaryRedirect, http.StatusPermanentRedirect} {
		t.Run(strconv.Itoa(status), func(t *testing.T) {
			client, capture, location := newRedirectTestAPIClient(
				t,
				status,
				"/nupf-ee/v1/ee-subscriptions",
			)
			request := &CreateSubscriptionRequest{}
			request.SetRequestBody(models.Upf_EvtExpos_CreateEventSubscription{
				Subscription: &models.Upf_EvtExpos_UpfEventSubscription{
					EventNotifyUri:      "https://consumer.example.com/notify",
					NotifyCorrelationId: "correlation-id",
				},
			})

			response, err := client.SubscriptionsCollectionApi.CreateSubscription(context.Background(), request)

			require.Nil(t, response)
			require.Error(t, err)
			var apiError openapi.GenericOpenAPIError
			require.ErrorAs(t, err, &apiError)
			require.Equal(t, status, apiError.ErrorStatus)
			errorModel, ok := apiError.Model().(CreateSubscriptionError)
			require.True(t, ok)
			require.Equal(t, location, errorModel.Location)
			require.Equal(t, redirectTargetNFID, errorModel.Var3gpp_Sbi_Target_Nf_Id)
			require.NotNil(t, errorModel.RedirectResponse)
			require.Equal(t, "TEMPORARY_REDIRECTION", errorModel.RedirectResponse.Cause)

			captured := capture.snapshot()
			require.Equal(t, int32(1), captured.originRequests)
			require.Zero(t, captured.targetRequests)
			require.Equal(t, int32(1), captured.originBodies)
			require.Equal(t, http.MethodPost, captured.originMethod)
			require.Equal(t, "/nupf-ee/v1/ee-subscriptions", captured.originPath)
			require.Contains(t, captured.originBody, `"subscription"`)
		})
	}
}
