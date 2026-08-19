package oauth

import (
	"context"
	"sync"
	"time"

	"golang.org/x/oauth2"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi/nrf/AccTok"
)

// cachedToken stores OAuth token with absolute expiry timestamp
type cachedToken struct {
	Response   models.Nrf_AccTok_AccessTokenRsp
	ExpiryTime int64 // absolute Unix timestamp when token expires
}

var tokenMap sync.Map
var clientMap sync.Map

func GetTokenCtx(
	nfType, targetNF models.Nrf_NFMgmt_NFType,
	nfId, nrfUri, scope string,
) (context.Context, *models.ProblemDetails, error) {
	tok, pd, err := sendAccTokenReq(nfType, targetNF, nfId, nrfUri, scope)
	if err != nil {
		return nil, pd, err
	}
	return context.WithValue(context.Background(),
		openapi.ContextOAuth2, tok), pd, nil
}

func sendAccTokenReq(
	nfType, targetNF models.Nrf_NFMgmt_NFType,
	nfId, nrfUri, scope string,
) (oauth2.TokenSource, *models.ProblemDetails, error) {
	var client *AccTok.APIClient

	if val, ok := clientMap.Load(nrfUri); ok {
		client = val.(*AccTok.APIClient)
	} else {
		configuration := AccTok.NewConfiguration()
		configuration.SetBasePath(nrfUri)
		client = AccTok.NewAPIClient(configuration)
		clientMap.Store(nrfUri, client)
	}

	// Check if we have a valid cached token
	if val, ok := tokenMap.Load(scope); ok {
		cached := val.(cachedToken)
		// Compare current time with absolute expiry timestamp
		if time.Now().Unix() < cached.ExpiryTime {
			token := &oauth2.Token{
				AccessToken: cached.Response.Access_token,
				TokenType:   cached.Response.Token_type,
				Expiry:      time.Unix(cached.ExpiryTime, 0),
			}
			return oauth2.StaticTokenSource(token), nil, nil
		}
	}

	req := &AccTok.AccessTokenRequestRequest{}
	req.SetGrantType("client_credentials")
	req.SetNfInstanceId(nfId)
	req.SetNfType(nfType)
	req.SetTargetNfType(targetNF)
	req.SetScope(scope)

	res, err := client.AccessTokenRequestApi.AccessTokenRequest(
		context.Background(), req)

	if err == nil {
		// Calculate absolute expiry time: current time + expires_in seconds
		expiryTime := time.Now().Unix() + int64(res.Nrf_AccTok_AccessTokenRsp.Expires_in)
		cached := cachedToken{
			Response:   *res.Nrf_AccTok_AccessTokenRsp,
			ExpiryTime: expiryTime,
		}
		tokenMap.Store(scope, cached)

		token := &oauth2.Token{
			AccessToken: res.Nrf_AccTok_AccessTokenRsp.Access_token,
			TokenType:   res.Nrf_AccTok_AccessTokenRsp.Token_type,
			Expiry:      time.Unix(expiryTime, 0),
		}
		return oauth2.StaticTokenSource(token), nil, nil
	} else {
		return nil, nil, openapi.ReportError("server no response")
	}
}
