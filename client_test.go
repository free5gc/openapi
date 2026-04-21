package openapi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParameterToString(t *testing.T) {
	var testCases = []struct {
		name   string
		obj    interface{}
		format string
		para   string
	}{
		{
			name:   "strings csv",
			obj:    []string{"red", "blue", "yellow"},
			format: "csv",
			para:   "red,blue,yellow",
		},
		{
			name:   "strings pipes",
			obj:    []string{"red", "blue", "yellow"},
			format: "pipes",
			para:   "red|blue|yellow",
		},
		{
			name:   "strings ssv",
			obj:    []string{"red", "blue", "yellow"},
			format: "ssv",
			para:   "red blue yellow",
		},
		{
			name:   "strings json",
			obj:    []string{"red", "blue", "yellow"},
			format: "application/json",
			para:   "[\"red\",\"blue\",\"yellow\"]",
		},
		{
			name: "obj json",
			obj: struct {
				Name   string `json:"name"`
				Age    int    `json:"age"`
				Active bool   `json:"active"`
			}{
				Name:   "free5GC",
				Age:    3,
				Active: true,
			},
			format: "application/json",
			para:   "{\"name\":\"free5GC\",\"age\":3,\"active\":true}",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			para := ParameterToString(tc.obj, tc.format)
			require.Equal(t, tc.para, para)
		})
	}
}

// TestMultipartDeserialize_MalformedBody pins the fix for free5gc/free5gc#1026.
// Before the fix, a body whose Content-Type declared multipart/related but
// whose payload was not a valid MIME multipart stream would panic on
// part.Header.Get at client.go:608 because mime/multipart.NextPart can return
// a nil *Part alongside a non-EOF error. The SMF then surfaced HTTP 500 on
// every malformed POST /nsmf-pdusession/v1/sm-contexts.
func TestMultipartDeserialize_MalformedBody(t *testing.T) {
	type dummy struct{}

	cases := []struct {
		name string
		body []byte
	}{
		{name: "plain text body", body: []byte("hello, not multipart")},
		{name: "empty body", body: []byte("")},
		{name: "json body", body: []byte(`{"foo":"bar"}`)},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var out dummy
			require.NotPanics(t, func() {
				err := MultipartDeserialize(tc.body, &out, "boundaryX")
				require.Error(t, err)
			})
		})
	}
}
