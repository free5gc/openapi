package openapi

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/openapi/models"
)

func TestSnssaiEqualFold(t *testing.T) {
	testCases := []struct {
		name     string
		s        models.Snssai
		t        models.Snssai
		expected bool
	}{
		{
			name:     "case insensitive 1",
			s:        models.Snssai{Sst: 1, Sd: "abcdef"},
			t:        models.Snssai{Sst: 1, Sd: "ABCDEF"},
			expected: true,
		},
		{
			name:     "case insensitive 2",
			s:        models.Snssai{Sst: 1, Sd: "012def"},
			t:        models.Snssai{Sst: 1, Sd: "012DeF"},
			expected: true,
		},
		{
			name:     "fail case",
			s:        models.Snssai{Sst: 1, Sd: "012def"},
			t:        models.Snssai{Sst: 1, Sd: "012Dee"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SnssaiEqualFold(tc.s, tc.t)
			require.Equal(t, tc.expected, result)
		})
	}
}

