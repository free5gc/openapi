package openapi

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"

	"github.com/free5gc/openapi/models"
)

type CCAClaims struct {
	Iat int32
	Exp int32
	jwt.StandardClaims
}

func GenerateClientCredentialAssertion(
	sub, aud, keyPath string,
) (string, error) {
	var expiration int32 = 1000
	now := int32(time.Now().Unix())

	accessTokenClaims := CCAClaims{
		Iat: now,
		Exp: now + expiration, // access_token is authorized for use
		StandardClaims: jwt.StandardClaims{
			Subject:  sub,
			Audience: aud,
		},
	}

	// Use RSA as a signing method
	signKey, err := ParsePrivateKeyFromPEM(keyPath)
	if err != nil {
		return "", errors.Wrapf(err, "gen CCAClaims")
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS512"), accessTokenClaims)
	accessToken, err := token.SignedString(signKey)
	if err != nil {
		return "", errors.Wrapf(err, "gen CCAClaims")
	}
	return accessToken, nil
}

func VerifyOAuth(
	authorization, serviceName, certPath string,
) error {
	verifyKey, err := ParsePublicKeyFromPEM(certPath)
	if err != nil {
		return errors.Wrapf(err, "verify OAuth")
	}

	auth_fields := strings.Fields(authorization)
	if len(auth_fields) < 2 {
		return errors.Errorf("verify OAuth Authorization header invalid")
	}

	access_token := auth_fields[1]
	token, err := jwt.ParseWithClaims(
		access_token,
		&models.AccessTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, errors.Wrapf(err, "Unexpected signing method")
			}
			if token.Header["alg"] != "RS512" {
				return nil, errors.Wrapf(err, "Unexpected signing method")
			}
			return verifyKey, nil
		})
	if err != nil {
		return errors.Wrapf(err, "verify OAuth parse")
	}

	if !verifyScope(token.Claims.(*models.AccessTokenClaims).Scope, serviceName) {
		return errors.Wrapf(err, "verify OAuth scope")
	}
	return nil
}

func verifyScope(scope, serviceName string) bool {
	if len(serviceName) == 0 {
		return true
	}
	if len(scope) != 0 {
		scopeSplit := strings.Fields(scope)
		found := false
		for _, item := range scopeSplit {
			if item == serviceName {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	} else {
		return false
	}
	return true
}

func GenerateRootCertificate(
	rootCertPath string,
	rootPrivKey *rsa.PrivateKey,
) (*x509.Certificate, error) {
	rootCert, err := GenerateCertificate(
		"", "", rootCertPath, &rootPrivKey.PublicKey, nil, rootPrivKey)
	if err != nil {
		return nil, errors.Wrapf(err, "gen root cert")
	}
	return rootCert, nil
}

func GenerateCertificate(
	nfType, nfId, certPemPath string,
	pubKey *rsa.PublicKey,
	rootCert *x509.Certificate,
	rootPrivKey *rsa.PrivateKey,
) (*x509.Certificate, error) {
	max := new(big.Int)
	max.Exp(big.NewInt(16), big.NewInt(40), nil)
	sn, _ := rand.Int(rand.Reader, max)

	temp := &x509.Certificate{
		SerialNumber: sn,
		Subject: pkix.Name{
			Country:            []string{"TW"},
			Province:           []string{"Taiwan"},
			Locality:           []string{"Hsinchu"},
			Organization:       []string{"free5gc"},
			OrganizationalUnit: []string{"free5gc"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(10, 0, 0),
	}

	if nfType != "" {
		temp.Subject.CommonName = strings.ToUpper(nfType)
		temp.DNSNames = []string{nfType}
	}
	if nfId != "" {
		uri, err := url.Parse("urn:uuid:" + nfId)
		if err != nil {
			return nil, errors.Wrapf(err, "gen cert url")
		}
		temp.URIs = []*url.URL{uri}
	}
	if rootCert == nil {
		// generate self-signed certificate
		rootCert = temp
	}

	b, err := x509.CreateCertificate(
		rand.Reader, temp, rootCert, pubKey, rootPrivKey)
	if err != nil {
		return nil, errors.Wrapf(err, "gen cert create")
	}

	cert, err := x509.ParseCertificate(b)
	if err != nil {
		return nil, errors.Wrapf(err, "gen cert parse")
	}

	if certPemPath != "" {
		out := &bytes.Buffer{}
		err = pem.Encode(out, &pem.Block{Type: "CERTIFICATE", Bytes: b})
		if err != nil {
			return nil, errors.Wrapf(err, "gen cert file encode")
		}

		certFile, err := os.Create(certPemPath)
		if err != nil {
			return nil, errors.Wrapf(err, "gen cert file create")
		}
		defer certFile.Close() // nolint

		_, err = out.WriteTo(certFile)
		if err != nil {
			return nil, errors.Wrapf(err, "gen cert file write")
		}
	}

	return cert, nil
}

func ParsePublicKeyFromPEM(pubPemPath string) (*rsa.PublicKey, error) {
	b, err := os.ReadFile(pubPemPath)
	if err != nil {
		return nil, errors.Wrapf(err, "pubkey read")
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, errors.Wrapf(err, "pubkey parse")
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
