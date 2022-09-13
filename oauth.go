package openapi

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
)

func CheckOAuth(authorization string, scopes []string) (bool, error) {
	key, err := parseNRFCert()
	if err != nil {
		return false, err
	}

	auth_header := strings.Fields(authorization)

	if len(auth_header) > 0 {
		access_token := auth_header[1]

		return VerifyAccessTokenAndScopes(access_token, key, scopes)
	} else {
		return false, fmt.Errorf("Authorization header invalid")
	}
}

func parseNRFCert() (*rsa.PublicKey, error) {
	f, err := os.ReadFile("./config/cert/nrf.pem")
	if err != nil {
		return nil, err
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(f)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

func VerifyAccessTokenAndScopes(access_token string, key *rsa.PublicKey, scopes []string) (bool, error) {
	if claims, err := VerifyAccessToken(access_token, key); err != nil {
		return false, err
	} else {
		return verifyScopes(claims, scopes)
	}
}

func verifyScopes(claims jwt.MapClaims, scopes []string) (bool, error) {
	if len(scopes) == 0 {
		return true, nil
	} else if claims != nil {
		if err := claims.Valid(); err != nil {
			return false, fmt.Errorf("Claims invalid %+v", err.Error())
		}
		for _, scope := range scopes {
			if claimScope, ok := claims["scope"]; !ok {
				return false, fmt.Errorf("JWT Claims did not contain a scope.")
			} else {
				claimScopesSplit := strings.Fields(claimScope.(string))
				found := false
				for _, item := range claimScopesSplit {
					if item == scope {
						found = true
						break
					}
				}
				if !found {
					return false, fmt.Errorf("Did not find scope in claims")
				}
			}
		}
		return true, nil
	} else {
		return false, fmt.Errorf("Claims was nil")
	}
}

func VerifyAccessToken(access_token string, key *rsa.PublicKey) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(access_token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		if token.Header["alg"] != "RS512" {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	})

	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			return claims, nil
		} else {
			return nil, fmt.Errorf("Token claims error %+v", err)
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, fmt.Errorf("Token is malformed %+v", ve)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, fmt.Errorf("Token is expired %+v", ve)
		} else {
			return nil, fmt.Errorf("Token had an unknown error %+v", ve)
		}
	} else {
		// Unknown error
		return nil, fmt.Errorf("An unknown error occurred validating the token. That's all we know")
	}
}

func CreateContext(oauth bool, nfId, NrfUri, reqNf string) context.Context {
	var c context.Context
	additional_params := url.Values{}
	additional_params.Add("OAuth", strconv.FormatBool(oauth))

	if oauth == true {
		additional_params.Add("grant_type", "client_credentials")
		additional_params.Add("nfInstanceId", nfId)
		additional_params.Add("NrfUri", NrfUri)
		additional_params.Add("nfType", reqNf)
	}

	c = context.WithValue(context.Background(), ContextOAuthAdditionalParams, additional_params)

	return c
}
