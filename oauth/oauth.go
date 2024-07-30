package oauth

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	"github.com/free5gc/openapi/models"
)

type CCAClaims struct {
	Iat int32
	Exp int32
	jwt.RegisteredClaims
}

func GenerateClientCredentialAssertion(
	sub, aud, keyPath string,
) (string, error) {
	var expiration int32 = 1000
	now := int32(time.Now().Unix())

	accessTokenClaims := CCAClaims{
		Iat: now,
		Exp: now + expiration, // access_token is authorized for use
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:  sub,
			Audience: jwt.ClaimStrings{aud},
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

func ParsePrivateKeyFromPEM(privPemPath string) (*rsa.PrivateKey, error) {
	b, err := os.ReadFile(privPemPath)
	if err != nil {
		return nil, errors.Wrapf(err, "privkey read")
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, errors.Wrapf(err, "privkey parse")
	}

	return privKey, nil
}

func ParseCertFromPEM(certPemPath string) (*x509.Certificate, error) {
	b, err := os.ReadFile(certPemPath)
	if err != nil {
		return nil, errors.Wrapf(err, "read cert pem")
	}

	block, _ := pem.Decode(b)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.Wrapf(err, "parse cert pem")
	}

	return cert, nil
}

func GenerateRSAKeyPair(pubPemPath, privPemPath string) (*rsa.PrivateKey, error) {
	const rsaKeyBitSize = 2048

	// generate key
	privKey, err := rsa.GenerateKey(rand.Reader, rsaKeyBitSize)
	if err != nil {
		return nil, errors.Wrapf(err, "generate rsa key")
	}

	if pubPemPath != "" {
		// dump public key to file
		pubPem, err := os.Create(pubPemPath)
		if err != nil {
			return nil, errors.Wrapf(err, "generate rsa pub key create file")
		}
		defer pubPem.Close() // nolint

		pubKey := &privKey.PublicKey
		pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
		if err != nil {
			return nil, errors.Wrapf(err, "generate rsa pub key marshal")
		}
		pubKeyBlock := &pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pubKeyBytes,
		}

		err = pem.Encode(pubPem, pubKeyBlock)
		if err != nil {
			return nil, errors.Wrapf(err, "generate rsa pub key pem")
		}
	}

	if privPemPath != "" {
		// dump private key to file
		privPem, err := os.Create(privPemPath)
		if err != nil {
			return nil, errors.Wrapf(err, "generate rsa priv key create file")
		}
		defer privPem.Close() // nolint

		privKeyBlock := &pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privKey),
		}

		err = pem.Encode(privPem, privKeyBlock)
		if err != nil {
			return nil, errors.Wrapf(err, "generate rsa priv key pem")
		}
	}
	return privKey, nil
}

func GetNFCertFileName(nfType, nfId string) string {
	if nfId != "" {
		return strings.ToLower(nfType) + "_" + nfId + ".pem"
	}
	return strings.ToLower(nfType) + ".pem"
}

func GetNFCertPath(base, nfType, nfId string) string {
	// Note: NF's cert should be put in the same base path
	return filepath.Join(base, GetNFCertFileName(nfType, nfId))
}
