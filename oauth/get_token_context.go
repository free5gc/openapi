package oauth

import (
	"context"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
)

func GetTokenCtx(
	oauth2Required bool, nfType models.NfType, nfId, nrfUri, scope, targetNF string,
) (context.Context, *models.ProblemDetails, error) {
	if !oauth2Required {
		return context.TODO(), nil, nil
	}
	tok, pd, err := SendAccTokenReq(nfId, nfType, scope, targetNF, nrfUri)
	if err != nil {
		return nil, pd, err
	}
	return context.WithValue(context.Background(),
		openapi.ContextOAuth2, tok), pd, nil
}
