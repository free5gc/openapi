package oauth

import (
	"testing"
	"time"

	"github.com/free5gc/openapi/models"
	"github.com/golang-jwt/jwt/v5"
)

func TestVerifyOAuthRejectsExpiredGeneratedModelToken(t *testing.T) {
	dir := t.TempDir()
	pubPath, privPath := dir+"/public.pem", dir+"/private.pem"
	key, err := GenerateRSAKeyPair(pubPath, privPath)
	if err != nil {
		t.Fatal(err)
	}
	claims := models.AccessTokenClaims{
		Scope: "svc",
		Exp:   int32(time.Now().Add(-time.Minute).Unix()),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "issuer",
		},
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	signed, err := tok.SignedString(key)
	if err != nil {
		t.Fatal(err)
	}
	if err := VerifyOAuth("Bearer "+signed, "svc", pubPath); err == nil {
		t.Fatal("expired token was accepted")
	}
}

func TestVerifyOAuthRequiresBearerScheme(t *testing.T) {
	dir := t.TempDir()
	pubPath, privPath := dir+"/public.pem", dir+"/private.pem"
	key, err := GenerateRSAKeyPair(pubPath, privPath)
	if err != nil {
		t.Fatal(err)
	}
	claims := models.AccessTokenClaims{
		Scope: "svc",
		Exp:   int32(time.Now().Add(time.Minute).Unix()),
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	signed, err := tok.SignedString(key)
	if err != nil {
		t.Fatal(err)
	}
	if err := VerifyOAuth("Basic "+signed, "svc", pubPath); err == nil {
		t.Fatal("non-Bearer authorization scheme was accepted")
	}
}

