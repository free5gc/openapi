package oauth

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
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

// TokenRequest identifies the consumer, target selector, NRF, and scope used
// to acquire and cache an access token. TargetNFInstanceID selects a specific
// producer; otherwise ConsumerNFType and TargetNFType form a type-level target.
type TokenRequest struct {
	ConsumerNFType       models.NrfNfManagementNfType
	ConsumerNFInstanceID string
	TargetNFType         models.NrfNfManagementNfType
	TargetNFInstanceID   string
	NRFURI               string
	Scope                string
}

func (request TokenRequest) validate() error {
	if isBlank(request.ConsumerNFInstanceID) {
		return errors.New("invalid token request: consumer NF instance ID is empty")
	}
	if isBlank(request.NRFURI) {
		return errors.New("invalid token request: NRF URI is empty")
	}
	if isBlank(request.Scope) {
		return errors.New("invalid token request: scope is empty")
	}

	if isBlank(request.TargetNFInstanceID) {
		if isBlank(string(request.ConsumerNFType)) {
			return errors.New("invalid type-level token request: consumer NF type is empty")
		}
		if isBlank(string(request.TargetNFType)) {
			return errors.New("invalid type-level token request: target NF type is empty")
		}
	}
	return nil
}

func isBlank(value string) bool {
	return strings.TrimSpace(value) == ""
}

type tokenCacheKey struct {
	ConsumerNFType       models.NrfNfManagementNfType
	ConsumerNFInstanceID string
	TargetNFType         models.NrfNfManagementNfType
	TargetNFInstanceID   string
	NRFURI               string
	Scope                string
}

var tokenMap sync.Map
var clientMap sync.Map

func GetTokenCtx(
<<<<<<< HEAD
	nfType, targetNF models.Nrf_NFMgmt_NFType,
	nfId, nrfUri, scope string,
=======
	request TokenRequest,
>>>>>>> d8cc967 (fix: add token request and audience validation)
) (context.Context, *models.ProblemDetails, error) {
	if err := request.validate(); err != nil {
		return nil, nil, err
	}
	tok, pd, err := sendAccTokenReq(request)
	if err != nil {
		return nil, pd, err
	}
	return context.WithValue(context.Background(),
		openapi.ContextOAuth2, tok), pd, nil
}

func sendAccTokenReq(
<<<<<<< HEAD
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
=======
	request TokenRequest,
) (oauth2.TokenSource, *models.ProblemDetails, error) {
	cacheKey := tokenCacheKey(request)
	var client *AccessToken.APIClient

	if val, ok := clientMap.Load(request.NRFURI); ok {
		client = val.(*AccessToken.APIClient)
	} else {
		configuration := AccessToken.NewConfiguration()
		configuration.SetBasePath(request.NRFURI)
		client = AccessToken.NewAPIClient(configuration)
		clientMap.Store(request.NRFURI, client)
>>>>>>> d8cc967 (fix: add token request and audience validation)
	}

	// Check if we have a valid cached token
	if val, ok := tokenMap.Load(cacheKey); ok {
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
	req.SetNfInstanceId(request.ConsumerNFInstanceID)
	req.SetScope(request.Scope)
	if !isBlank(string(request.ConsumerNFType)) {
		req.SetNfType(request.ConsumerNFType)
	}
	if !isBlank(string(request.TargetNFType)) {
		req.SetTargetNfType(request.TargetNFType)
	}
	if !isBlank(request.TargetNFInstanceID) {
		req.SetTargetNfInstanceId(request.TargetNFInstanceID)
	}

	res, err := client.AccessTokenRequestApi.AccessTokenRequest(
		context.Background(), req)

	if err == nil {
		// Calculate absolute expiry time: current time + expires_in seconds
		expiryTime := time.Now().Unix() + int64(res.Nrf_AccTok_AccessTokenRsp.Expires_in)
		cached := cachedToken{
			Response:   *res.Nrf_AccTok_AccessTokenRsp,
			ExpiryTime: expiryTime,
		}
		tokenMap.Store(cacheKey, cached)

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
