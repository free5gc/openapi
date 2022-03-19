package openapi

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)

type CCAClaims struct {
	Iat int32
	Exp int32
	jwt.StandardClaims
}

func GenClientCredentialAssertion(sub string, aud string, keyPath string) (string, error) {
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
	signBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return "", err
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS512"), accessTokenClaims)
	accessToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func OAuthVerify(request *httpwrapper.Request, serviceName string, pubKeyPath string) *httpwrapper.Response {
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		fmt.Println("VerifyBytes failed")
		return httpwrapper.NewResponse(http.StatusBadRequest, nil, err)
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		fmt.Println("VerifyKey failed")
		return httpwrapper.NewResponse(http.StatusBadRequest, nil, err)
	}
	if request.Header.Get("Authorization") == "" {
		fmt.Println("no access token")
		return httpwrapper.NewResponse(http.StatusBadRequest, nil, err)
	}
	auth := strings.Split(request.Header.Get("Authorization"), " ")
	tokenString := strings.TrimSpace(auth[1])
	token, err := jwt.ParseWithClaims(tokenString, &models.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return httpwrapper.NewResponse(http.StatusUnauthorized, nil, err)
	}

	if !VerifyScope(token.Claims.(*models.AccessTokenClaims).Scope, serviceName) {
		return httpwrapper.NewResponse(http.StatusForbidden, nil, nil)
	}
	return nil
}

func VerifyScope(scope string, serviceName string) bool {
	return strings.Contains(scope, serviceName)
}

func LoadRootKey(rootKeyPath string) (*rsa.PrivateKey, error) {
	priv, err := ioutil.ReadFile(rootKeyPath)
	if err != nil {
		fmt.Println("Read root key file error: ", err)
	}
	privPem, _ := pem.Decode(priv)
	privPemBytes := privPem.Bytes
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(privPemBytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(privPemBytes); err != nil {
			fmt.Println("parsed key error: ", err)
		}
	}
	var privateKey *rsa.PrivateKey
	var ok bool
	privateKey, ok = parsedKey.(*rsa.PrivateKey)
	if !ok {
		fmt.Println("no ok")
	}

	return privateKey, nil
}

func GenCertificate(nfType string, nfId string, rootPemPath string,
	rootKeyPath string, pemPath string, keyPath string) error {
	max := new(big.Int)
	max.Exp(big.NewInt(16), big.NewInt(40), nil)
	sn, _ := rand.Int(rand.Reader, max)
	uri, err := url.Parse("urn:uuid:" + nfId)
	if err != nil {
		fmt.Println("Parse url error: ", err)
	}

	temp := &x509.Certificate{
		SerialNumber: sn,
		Subject: pkix.Name{
			Country:            []string{"TW"},
			Province:           []string{"Taiwan"},
			Locality:           []string{"Hsinchu"},
			Organization:       []string{"free5gc"},
			OrganizationalUnit: []string{"free5gc"},
			CommonName:         strings.ToUpper(nfType),
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(1, 0, 0),
		DNSNames:  []string{nfType},
		URIs:      []*url.URL{uri},
	}

	keyString, err := ioutil.ReadFile(keyPath)
	if err != nil {
		fmt.Println("Read key file error: ", err)
		return err
	}
	keyBlock, _ := pem.Decode([]byte(keyString))
	certPrivKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		fmt.Println("Parse private key error: ", err)
		return err
	}

	if err != nil {
		fmt.Println("Generate Key error: ", err)
		return err
	}

	rootCrt, err := ioutil.ReadFile(rootPemPath)
	if err != nil {
		fmt.Println("Read root certificate error: ", err)
		return err
	}
	block, _ := pem.Decode([]byte(rootCrt))
	ca, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("Parse root certificate error: ", err)
		return err
	}

	rootKey, _ := LoadRootKey(rootKeyPath)

	cert, err := x509.CreateCertificate(rand.Reader, temp, ca, &certPrivKey.PublicKey, rootKey)
	if err != nil {
		fmt.Println("Create Certificate error: ", err)
		return err
	}
	out := &bytes.Buffer{}
	pem.Encode(out, &pem.Block{Type: "CERTIFICATE", Bytes: cert})
	certFile, _ := os.OpenFile(pemPath, os.O_RDWR|os.O_CREATE, 0666)
	out.WriteTo(certFile)
	certFile.Close()
	fmt.Println("finish generate cert")

	return nil
}
