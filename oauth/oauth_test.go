package oauth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/free5gc/openapi/models"
)

func TestVerifyOAuthClaims(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate test key: %v", err)
	}

	publicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		t.Fatalf("marshal test public key: %v", err)
	}
	certPath := filepath.Join(t.TempDir(), "nrf-public.pem")
	if err := os.WriteFile(certPath, pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKey,
	}), 0600); err != nil {
		t.Fatalf("write test public key: %v", err)
	}

	const (
		issuer     = "nrf-instance-id"
		instanceID = "smf-instance-id"
		service    = "nsmf-pdusession"
	)
	now := time.Now()
	validClaims := func() accessTokenClaims {
		return accessTokenClaims{
			Scope: service,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    issuer,
				Audience:  jwt.ClaimStrings{instanceID},
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
			},
		}
	}

	tests := []struct {
		name    string
		change  func(*accessTokenClaims)
		wantErr bool
	}{
		{name: "valid token"},
		{
			name: "wrong issuer",
			change: func(claims *accessTokenClaims) {
				claims.Issuer = "another-nrf"
			},
			wantErr: true,
		},
		{
			name: "wrong audience",
			change: func(claims *accessTokenClaims) {
				claims.Audience = jwt.ClaimStrings{"another-smf"}
			},
			wantErr: true,
		},
		{
			name: "expired token",
			change: func(claims *accessTokenClaims) {
				claims.ExpiresAt = jwt.NewNumericDate(now.Add(-time.Minute))
			},
			wantErr: true,
		},
		{
			name: "missing expiration",
			change: func(claims *accessTokenClaims) {
				claims.ExpiresAt = nil
			},
			wantErr: true,
		},
		{
			name: "wrong scope",
			change: func(claims *accessTokenClaims) {
				claims.Scope = "namf-comm"
			},
			wantErr: true,
		},
	}

	policy := AudiencePolicy{
		NFInstanceID: instanceID,
		NFType:       models.Nrf_NFMgmt_NFType_SMF,
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			claims := validClaims()
			if test.change != nil {
				test.change(&claims)
			}
			token, err := jwt.NewWithClaims(jwt.SigningMethodRS512, claims).SignedString(privateKey)
			if err != nil {
				t.Fatalf("sign test token: %v", err)
			}

			err = VerifyOAuth("Bearer "+token, service, policy, issuer, certPath)
			if (err != nil) != test.wantErr {
				t.Fatalf("VerifyOAuth() error = %v, want error = %t", err, test.wantErr)
			}
		})
	}
}

func TestAccessTokenClaimsExpirationRoundTrip(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate test key: %v", err)
	}

	now := time.Now()
	for _, test := range []struct {
		name      string
		expiresAt time.Time
		wantErr   bool
	}{
		{name: "valid", expiresAt: now.Add(time.Hour)},
		{name: "expired", expiresAt: now.Add(-time.Minute), wantErr: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			claims := models.Nrf_AccTok_AccessTokenClaims{
				Scope: "nsmf-pdusession",
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "nrf-instance-id",
					Subject:   "consumer-instance-id",
					Audience:  jwt.ClaimStrings{"smf-instance-id"},
					ExpiresAt: jwt.NewNumericDate(test.expiresAt),
				},
			}
			payload, err := json.Marshal(claims)
			if err != nil {
				t.Fatalf("marshal claims: %v", err)
			}
			var decoded models.Nrf_AccTok_AccessTokenClaims
			if err := json.Unmarshal(payload, &decoded); err != nil {
				t.Fatalf("unmarshal claims: %v", err)
			}
			if decoded.ExpiresAt == nil || decoded.ExpiresAt.Unix() != test.expiresAt.Unix() {
				t.Fatalf("expiration lost on JSON round trip: %s", payload)
			}
			if decoded.Issuer != claims.Issuer || decoded.Subject != claims.Subject ||
				len(decoded.Audience) != 1 || decoded.Audience[0] != claims.Audience[0] {
				t.Fatalf("standard identity claims lost on JSON round trip: %s", payload)
			}

			token, err := jwt.NewWithClaims(jwt.SigningMethodRS512, claims).SignedString(privateKey)
			if err != nil {
				t.Fatalf("sign test token: %v", err)
			}
			_, err = jwt.ParseWithClaims(token, &models.Nrf_AccTok_AccessTokenClaims{},
				func(_ *jwt.Token) (interface{}, error) { return &privateKey.PublicKey, nil },
				jwt.WithValidMethods([]string{"RS512"}), jwt.WithExpirationRequired())
			if (err != nil) != test.wantErr {
				t.Fatalf("parse token error = %v, want error = %t", err, test.wantErr)
			}
		})
	}
}

func TestGenerateClientCredentialAssertionStandardTimestamps(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate test key: %v", err)
	}
	keyPath := filepath.Join(t.TempDir(), "consumer-private.pem")
	if err := os.WriteFile(keyPath, pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}), 0600); err != nil {
		t.Fatalf("write test private key: %v", err)
	}

	assertion, err := GenerateClientCredentialAssertion("consumer-instance-id", "NRF", keyPath)
	if err != nil {
		t.Fatalf("generate CCA: %v", err)
	}
	parsed, err := jwt.ParseWithClaims(assertion, jwt.MapClaims{},
		func(_ *jwt.Token) (interface{}, error) { return &privateKey.PublicKey, nil },
		jwt.WithValidMethods([]string{"RS512"}), jwt.WithExpirationRequired())
	if err != nil {
		t.Fatalf("parse CCA: %v", err)
	}
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatalf("unexpected CCA claims type %T", parsed.Claims)
	}
	for _, name := range []string{"iat", "exp"} {
		if _, exists := claims[name]; !exists {
			t.Errorf("CCA missing standard %q claim", name)
		}
	}
	for _, name := range []string{"Iat", "Exp"} {
		if _, exists := claims[name]; exists {
			t.Errorf("CCA contains non-standard %q claim", name)
		}
	}
	iat, iatOK := claims["iat"].(float64)
	exp, expOK := claims["exp"].(float64)
	if !iatOK || !expOK || exp-iat != 1000 {
		t.Errorf("CCA timestamp claims = iat %v, exp %v; want 1000-second lifetime", claims["iat"], claims["exp"])
	}
}
