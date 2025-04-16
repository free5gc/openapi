package openapi

import (
	"testing"

	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestApiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		srvName  models.ServiceName
		expected string
	}{
		{
			name:     "v1",
			srvName:  models.ServiceName_NUPF_OAM,
			expected: "v1",
		},
		{
			name:     "v2",
			srvName:  models.ServiceName_NUDR_DR,
			expected: "v2",
		},
		{
			name:     "null",
			srvName:  models.ServiceName_NNSSAAF_NSSAA,
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ApiVersion(tc.srvName)
			require.Equal(t, tc.expected, result)
		})
	}
}

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

func TestExtSnssaiEqualFold(t *testing.T) {
	testCases := []struct {
		name     string
		s        models.ExtSnssai
		t        models.ExtSnssai
		expected bool
	}{
		{
			name:     "case insensitive 1",
			s:        models.ExtSnssai{Sst: 1, Sd: "abcdef"},
			t:        models.ExtSnssai{Sst: 1, Sd: "ABCDEF"},
			expected: true,
		},
		{
			name:     "case insensitive 2",
			s:        models.ExtSnssai{Sst: 1, Sd: "012def"},
			t:        models.ExtSnssai{Sst: 1, Sd: "012DeF"},
			expected: true,
		},
		{
			name:     "fail case",
			s:        models.ExtSnssai{Sst: 1, Sd: "012def"},
			t:        models.ExtSnssai{Sst: 1, Sd: "012Dee"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ExtSnssaiEqualFold(tc.s, tc.t)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestSnssaiHexToModels(t *testing.T) {
	testCases := []struct {
		name           string
		hexStr         string
		expectedSnssai *models.Snssai
		expectedErr    bool
	}{
		{
			name:           "pass case 1",
			hexStr:         "01abcdef",
			expectedSnssai: &models.Snssai{Sst: 1, Sd: "abcdef"},
			expectedErr:    false,
		},
		{
			name:           "pass case 2",
			hexStr:         "01",
			expectedSnssai: &models.Snssai{Sst: 1},
			expectedErr:    false,
		},
		{
			name:           "fail case 1",
			hexStr:         "1",
			expectedSnssai: nil,
			expectedErr:    true,
		},
		{
			name:           "fail case 2",
			hexStr:         "gh",
			expectedSnssai: nil,
			expectedErr:    true,
		},
		{
			name:           "fail case 3",
			hexStr:         "abcdef",
			expectedSnssai: nil,
			expectedErr:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := SnssaiHexToModels(tc.hexStr)
			if tc.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expectedSnssai, result)
		})
	}
}

func TestSnssaiModelsToHex(t *testing.T) {
	testCases := []struct {
		name        string
		snssai      models.Snssai
		expectedhex string
	}{
		{
			name:        "pass case 1",
			snssai:      models.Snssai{Sst: 1, Sd: "abcdef"},
			expectedhex: "01abcdef",
		},
		{
			name:        "pass case 2",
			snssai:      models.Snssai{Sst: 1},
			expectedhex: "01",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SnssaiModelsToHex(tc.snssai)
			require.Equal(t, tc.expectedhex, result)
		})
	}
}

func TestPlmnIdJsonToModels(t *testing.T) {
	testCases := []struct {
		name           string
		jsonStr        string
		expectedPlmnId *models.PlmnId
		expectedErr    bool
	}{
		{
			name:           "pass case",
			jsonStr:        "{\"mcc\":\"466\",\"mnc\":\"11\"}",
			expectedPlmnId: &models.PlmnId{Mcc: "466", Mnc: "11"},
			expectedErr:    false,
		},
		{
			name:           "fail case",
			jsonStr:        "46611",
			expectedPlmnId: nil,
			expectedErr:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := PlmnIdJsonToModels([]byte(tc.jsonStr))
			if tc.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expectedPlmnId, result)
		})
	}
}
