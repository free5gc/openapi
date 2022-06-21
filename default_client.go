//+build !debug

package openapi

import (
	"crypto/tls"
	"net/http"

	"sync"

	"golang.org/x/net/http2"
)

var initDefaultHttpClient sync.Once

// This function return the default HTTP client.
func GetDefaultHttpClient() *http.Client {
	initDefaultHttpClient.Do(func() {
		http.DefaultClient.Transport = &http2.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	})
	return http.DefaultClient
}
