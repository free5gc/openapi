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
