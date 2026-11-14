package AccTok

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/free5gc/openapi/models"
)

func TestAccessTokenRequestCCAIsRequestScoped(t *testing.T) {
	const header = "3gpp-Sbi-Client-Credentials"
	type received struct {
		assertion, consumer, target string
		inBody                      bool
	}
	requests := make(chan received, 3)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Error(err)
		}
		_, inBody := r.PostForm["clientCredentialsAssertion"]
		requests <- received{r.Header.Get(header), r.PostForm.Get("nfInstanceId"), r.PostForm.Get("targetNfInstanceId"), inBody}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"test-token","token_type":"Bearer","expires_in":60,"scope":"nausf-auth"}`))
	}))
	defer server.Close()
	cfg := NewConfiguration()
	cfg.SetBasePath(server.URL)
	cfg.SetHTTPClient(server.Client())
	client := NewAPIClient(cfg)
	for _, assertion := range []string{"consumer-a-assertion", "consumer-b-assertion", ""} {
		req := &AccessTokenRequestRequest{}
		req.SetGrantType("client_credentials")
		req.SetNfInstanceId("consumer-" + assertion)
		req.SetNfType(models.Nrf_NFMgmt_NFType_AMF)
		req.SetTargetNfType(models.Nrf_NFMgmt_NFType_AUSF)
		req.SetTargetNfInstanceId("producer-instance")
		req.SetScope("nausf-auth")
		if assertion != "" {
			req.SetClientCredentialsAssertion(assertion)
		}
		if _, err := client.AccessTokenRequestApi.AccessTokenRequest(context.Background(), req); err != nil {
			t.Fatal(err)
		}
		got := <-requests
		if got.assertion != assertion || got.consumer != "consumer-"+assertion || got.target != "producer-instance" || got.inBody {
			t.Fatalf("unexpected token request: %+v", got)
		}
	}
}
