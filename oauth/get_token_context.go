package oauth

import (
	"context"
	"sync"
	"time"

	"golang.org/x/oauth2"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi/nrf/AccessToken"
)

var tokenMap sync.Map
var clientMap sync.Map

func GetTokenCtx(
	nfType, targetNF models.NrfNfManagementNfType,
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
	nfType, targetNF models.NrfNfManagementNfType,
	nfId, nrfUri, scope string,
) (oauth2.TokenSource, *models.ProblemDetails, error) {
	var client *AccessToken.APIClient

	configuration := AccessToken.NewConfiguration()
	configuration.SetBasePath(nrfUri)

	if val, ok := clientMap.Load(configuration); ok {
		client = val.(*AccessToken.APIClient)
	} else {
		client = AccessToken.NewAPIClient(configuration)
		clientMap.Store(configuration, client)
	}

	var tok models.NrfAccessTokenAccessTokenRsp
	if val, ok := tokenMap.Load(scope); ok {
		tok = val.(models.NrfAccessTokenAccessTokenRsp)
		if int32(time.Now().Unix()) < tok.ExpiresIn {
			token := &oauth2.Token{
				AccessToken: tok.AccessToken,
				TokenType:   tok.TokenType,
				Expiry:      time.Unix(int64(tok.ExpiresIn), 0),
			}
			return oauth2.StaticTokenSource(token), nil, nil
		}
	}

	req := &AccessToken.AccessTokenRequestRequest{}
	req.SetGrantType("client_credentials")
	req.SetNfInstanceId(nfId)
	req.SetNfType(nfType)
	req.SetTargetNfType(targetNF)

	res, err := client.AccessTokenRequestApi.AccessTokenRequest(
		context.Background(), req)

	if err == nil {
		tokenMap.Store(scope, res.NrfAccessTokenAccessTokenRsp)
		token := &oauth2.Token{
			AccessToken: tok.AccessToken,
			TokenType:   tok.TokenType,
			Expiry:      time.Unix(int64(tok.ExpiresIn), 0),
		}
		return oauth2.StaticTokenSource(token), nil, nil
	} else {
		return nil, nil, openapi.ReportError("server no response")
	}
}
