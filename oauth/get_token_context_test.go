package oauth

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"golang.org/x/oauth2"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi/nrf/AccTok"
)

func TestGetTokenCtxRejectsInvalidSuccessResponse(t *testing.T) {
	for _, tc := range []struct{ name, body string }{
		{"missing token", `{"token_type":"Bearer","expires_in":60}`},
		{"empty token", `{"access_token":"","token_type":"Bearer","expires_in":60}`},
		{"blank token", `{"access_token":" \t ","token_type":"Bearer","expires_in":60}`},
		{"null token", `{"access_token":null,"token_type":"Bearer","expires_in":60}`},
		{"null response", `null`},
		{"empty object", `{}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var calls atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if calls.Add(1) == 1 {
					_, _ = w.Write([]byte(tc.body))
					return
				}
				_, _ = w.Write([]byte(`{"access_token":"valid-token","token_type":"Bearer","expires_in":60}`))
			}))
			defer server.Close()
			cfg := AccTok.NewConfiguration()
			cfg.SetBasePath(server.URL)
			cfg.SetHTTPClient(server.Client())
			clientMap.Store(server.URL, AccTok.NewAPIClient(cfg))
			request := TokenRequest{
				ConsumerNFInstanceID: "consumer", ConsumerNFType: models.Nrf_NFMgmt_NFType_AUSF,
				TargetNFType: models.Nrf_NFMgmt_NFType_UDM, NRFURI: server.URL, Scope: "nudm-ueau",
			}
			key := tokenCacheKey(request)
			t.Cleanup(func() { clientMap.Delete(server.URL); tokenMap.Delete(key) })
			ctx, _, err := GetTokenCtx(request)
			if err == nil || ctx != nil {
				t.Fatalf("invalid HTTP 200 accepted: error=%v, contextPresent=%t", err, ctx != nil)
			}
			if _, exists := tokenMap.Load(key); exists {
				t.Fatal("invalid token was cached")
			}
			// Retry must reach NRF; the following valid response may then be cached.
			for range 2 {
				ctx, _, err = GetTokenCtx(request)
				if err != nil {
					t.Fatal(err)
				}
				source, ok := ctx.Value(openapi.ContextOAuth2).(oauth2.TokenSource)
				if !ok {
					t.Fatal("missing token source")
				}
				token, err := source.Token()
				if err != nil {
					t.Fatal(err)
				}
				outbound, err := http.NewRequest(http.MethodGet, "http://producer.example/", nil)
				if err != nil {
					t.Fatal(err)
				}
				token.SetAuthHeader(outbound)
				if outbound.Header.Get("Authorization") != "Bearer valid-token" {
					t.Fatal("invalid authorization header")
				}
			}
			if calls.Load() != 2 {
				t.Fatalf("NRF requests = %d, want 2", calls.Load())
			}
		})
	}
}
