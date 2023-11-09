package send_util

import (
	"context"
	"sync"
	"time"

	"github.com/antihax/optional"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/Nnrf_AccessToken"
	"github.com/free5gc/openapi/models"
	"golang.org/x/oauth2"
)

var tokenMap sync.Map
var clientMap sync.Map

func SendAccTokenReq(
	nfId string, nfType models.NfType, scope, targetNF string, nrfUri string,
) (oauth2.TokenSource, *models.ProblemDetails, error) {

	var client *Nnrf_AccessToken.APIClient

	configuration := Nnrf_AccessToken.NewConfiguration()
	configuration.SetBasePath(nrfUri)

	if val, ok := clientMap.Load(configuration); ok {
		client = val.(*Nnrf_AccessToken.APIClient)
	} else {
		client = Nnrf_AccessToken.NewAPIClient(configuration)
		clientMap.Store(configuration, client)
	}

	var tok models.AccessTokenRsp

	if val, ok := tokenMap.Load(scope); ok {
		tok = val.(models.AccessTokenRsp)
		if int32(time.Now().Unix()) < tok.ExpiresIn {
			token := &oauth2.Token{
				AccessToken: tok.AccessToken,
				TokenType:   tok.TokenType,
				Expiry:      time.Unix(int64(tok.ExpiresIn), 0),
			}
			return oauth2.StaticTokenSource(token), nil, nil
		}
	}
	tok, res, err := client.AccessTokenRequestApi.AccessTokenRequest(context.Background(), "client_credentials",
		nfId, scope, &Nnrf_AccessToken.AccessTokenRequestParamOpts{
			NfType:       optional.NewInterface(nfType),
			TargetNfType: optional.NewInterface(targetNF),
		})

	if err == nil {
		tokenMap.Store(scope, tok)
		token := &oauth2.Token{
			AccessToken: tok.AccessToken,
			TokenType:   tok.TokenType,
			Expiry:      time.Unix(int64(tok.ExpiresIn), 0),
		}
		return oauth2.StaticTokenSource(token), nil, err
	} else if res != nil {
		if res.Status != err.Error() {
			return nil, nil, err
		}
		accesstoken_err := err.(openapi.GenericOpenAPIError).Model().(models.AccessTokenErr)
		var pd models.ProblemDetails
		pd.AccessTokenError = &accesstoken_err
		return nil, &pd, err
	} else {
		return nil, nil, openapi.ReportError("server no response")
	}
}
