package EvtExpos

import (
	"context"
	"net/http"
	"strconv"
	"testing"

	"github.com/free5gc/openapi"
	"github.com/stretchr/testify/require"
)

func TestDeleteSubscriptionRejectsRedirectWithoutReplay(t *testing.T) {
	for _, status := range []int{http.StatusTemporaryRedirect, http.StatusPermanentRedirect} {
		t.Run(strconv.Itoa(status), func(t *testing.T) {
			client, capture, location := newRedirectTestAPIClient(
				t,
				status,
				"/nupf-ee/v1/ee-subscriptions/sub-1",
			)
			request := &DeleteSubscriptionRequest{}
			request.SetSubscriptionId("sub-1")

			response, err := client.DefaultApi.DeleteSubscription(context.Background(), request)

			require.Nil(t, response)
			require.Error(t, err)
			var apiError openapi.GenericOpenAPIError
			require.ErrorAs(t, err, &apiError)
			require.Equal(t, status, apiError.ErrorStatus)
			errorModel, ok := apiError.Model().(DeleteSubscriptionError)
			require.True(t, ok)
			require.Equal(t, location, errorModel.Location)
			require.Equal(t, redirectTargetNFID, errorModel.Var3gpp_Sbi_Target_Nf_Id)
			require.NotNil(t, errorModel.RedirectResponse)
			require.Equal(t, "TEMPORARY_REDIRECTION", errorModel.RedirectResponse.Cause)

			captured := capture.snapshot()
			require.Equal(t, int32(1), captured.originRequests)
			require.Zero(t, captured.targetRequests)
			require.Zero(t, captured.originBodies)
			require.Equal(t, http.MethodDelete, captured.originMethod)
			require.Equal(t, "/nupf-ee/v1/ee-subscriptions/sub-1", captured.originPath)
			require.Empty(t, captured.originBody)
		})
	}
}
