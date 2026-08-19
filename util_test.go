package openapi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/openapi/models"
)

func TestApiVersion(t *testing.T) {
	testCases := []struct {
		name     string
		srvName  models.Nrf_NFMgmt_ServiceName
		expected string
	}{
		{
			name:     "v1",
			srvName:  models.Nrf_NFMgmt_ServiceName_NUPF_OAM,
			expected: "v1",
		},
		{
			name:     "v2",
			srvName:  models.Nrf_NFMgmt_ServiceName_NUDR_DR,
			expected: "v2",
		},
		{
			name:     "null",
			srvName:  models.Nrf_NFMgmt_ServiceName_NNSSAAF_NSSAA,
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
			name:     "case sd empty 1",
			s:        models.Snssai{Sst: 1, Sd: ""},
			t:        models.Snssai{Sst: 1, Sd: ""},
			expected: true,
		},
		{
			name:     "case sd empty 2",
			s:        models.Snssai{Sst: 1, Sd: ""},
			t:        models.Snssai{Sst: 1, Sd: "ffffff"},
			expected: true,
		},
		{
			name:     "case sd empty 3",
			s:        models.Snssai{Sst: 1, Sd: "ffffff"},
			t:        models.Snssai{Sst: 1, Sd: ""},
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
			name:     "case sd empty 1",
			s:        models.ExtSnssai{Sst: 1, Sd: ""},
			t:        models.ExtSnssai{Sst: 1, Sd: ""},
			expected: true,
		},
		{
			name:     "case sd empty 2",
			s:        models.ExtSnssai{Sst: 1, Sd: ""},
			t:        models.ExtSnssai{Sst: 1, Sd: "ffffff"},
			expected: true,
		},
		{
			name:     "case sd empty 3",
			s:        models.ExtSnssai{Sst: 1, Sd: "ffffff"},
			t:        models.ExtSnssai{Sst: 1, Sd: ""},
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

func TestPlmnIdEqualFold(t *testing.T) {
	type input struct {
		plmnId1 models.PlmnId
		plmnId2 models.PlmnId
	}
	tcs := []struct {
		name   string
		input  input
		expect bool
	}{
		{
			name: "unexpected(=diff): plmn1=empty",
			input: input{
				plmnId1: models.PlmnId{},
				plmnId2: models.PlmnId{
					Mcc: "208",
					Mnc: "93",
				},
			},
			expect: false,
		},
		{
			name: "unexpected(=diff): plmn2=empty",
			input: input{
				plmnId1: models.PlmnId{
					Mcc: "208",
					Mnc: "93",
				},
				plmnId2: models.PlmnId{},
			},
			expect: false,
		},
		{
			name: "unexpected but treated as same: plmn1=empty, plmn2=empty",
			input: input{
				plmnId1: models.PlmnId{},
				plmnId2: models.PlmnId{},
			},
			expect: true,
		},
		{
			name: "same: plmn1=20893=plmn2",
			input: input{
				plmnId1: models.PlmnId{
					Mcc: "208",
					Mnc: "93",
				},
				plmnId2: models.PlmnId{
					Mcc: "208",
					Mnc: "93",
				},
			},
			expect: true,
		},

		{
			name: "diff: plmn1=20893, plmn2:20894",
			input: input{
				plmnId1: models.PlmnId{
					Mcc: "208",
					Mnc: "93",
				},
				plmnId2: models.PlmnId{
					Mcc: "208",
					Mnc: "94",
				},
			},
			expect: false,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.name, func(t *testing.T) {
			same := PlmnIdEqual(tc.input.plmnId1, tc.input.plmnId2)
			require.Equal(t, tc.expect, same)
		})
	}
}

func TestPlmnIdNidEqualFold(t *testing.T) {
	type input struct {
		plmnIdNid1 models.PlmnIdNid
		plmnIdNid2 models.PlmnIdNid
	}
	tcs := []struct {
		name   string
		input  input
		expect bool
	}{
		{
			name: "unexpected(=diff): plmn1=empty",
			input: input{
				plmnIdNid1: models.PlmnIdNid{},
				plmnIdNid2: models.PlmnIdNid{
					Mcc: "208",
					Mnc: "93",
					Nid: "10000000001",
				},
			},
			expect: false,
		},
		{
			name: "unexpected(=diff): plmn2=empty",
			input: input{
				plmnIdNid1: models.PlmnIdNid{
					Mcc: "208",
					Mnc: "93",
					Nid: "10000000001",
				},
				plmnIdNid2: models.PlmnIdNid{},
			},
			expect: false,
		},
		{
			name: "unexpected but treated as same: plmn1=plmn2=empty",
			input: input{
				plmnIdNid1: models.PlmnIdNid{},
				plmnIdNid2: models.PlmnIdNid{},
			},
			expect: true,
		},
		{
			name: "same: plmn1=20893=plmn2",
			input: input{
				plmnIdNid1: models.PlmnIdNid{
					Mcc: "208",
					Mnc: "93",
				},
				plmnIdNid2: models.PlmnIdNid{
					Mcc: "208",
					Mnc: "93",
				},
			},
			expect: true,
		},
		{
			name: "same: nid1=20893-1abcdef0001=nid2",
			input: input{
				plmnIdNid1: models.PlmnIdNid{
					Mcc: "208",
					Mnc: "93",
					Nid: "1abcdef0001",
				},
				plmnIdNid2: models.PlmnIdNid{
					Mcc: "208",
					Mnc: "93",
					Nid: "1ABCDEF0001",
				},
			},
			expect: true,
		},
		{
			name: "diff: plmn1=20893, plmn2:20894",
			input: input{
				plmnIdNid1: models.PlmnIdNid{
					Mcc: "208",
					Mnc: "93",
				},
				plmnIdNid2: models.PlmnIdNid{
					Mcc: "208",
					Mnc: "94",
				},
			},
			expect: false,
		},
		{
			name: "diff: nid1=20893-10000000001, nid2=20893-10000000002",
			input: input{
				plmnIdNid1: models.PlmnIdNid{
					Mcc: "208",
					Mnc: "93",
					Nid: "10000000001",
				},
				plmnIdNid2: models.PlmnIdNid{
					Mcc: "208",
					Mnc: "93",
					Nid: "10000000002",
				},
			},
			expect: false,
		},
		{
			name: "diff: nid1=20893-10000000001, nid2=20894-10000000001",
			input: input{
				plmnIdNid1: models.PlmnIdNid{
					Mcc: "208",
					Mnc: "93",
					Nid: "10000000001",
				},
				plmnIdNid2: models.PlmnIdNid{
					Mcc: "208",
					Mnc: "94",
					Nid: "10000000001",
				},
			},
			expect: false,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.name, func(t *testing.T) {
			same := PlmnIdNidEqualFold(tc.input.plmnIdNid1, tc.input.plmnIdNid2)
			require.Equal(t, tc.expect, same)
		})
	}
}

func TestTaiEqualFold(t *testing.T) {
	type input struct {
		tai1 models.Tai
		tai2 models.Tai
	}
	tcs := []struct {
		name   string
		input  input
		expect bool
	}{
		{
			name: "unexpected(=diff): tai1=empty",
			input: input{
				tai1: models.Tai{},
				tai2: models.Tai{
					PlmnId: &models.PlmnId{
						Mcc: "208",
						Mnc: "93",
					},
					Tac: "000001",
				},
			},
			expect: false,
		},
		{
			name: "unexpected(=diff): tai2=empty",
			input: input{
				tai1: models.Tai{
					PlmnId: &models.PlmnId{
						Mcc: "208",
						Mnc: "93",
					},
					Tac: "000001",
				},
				tai2: models.Tai{},
			},
			expect: false,
		},
		{
			name: "unexpected(=diff): tai1=empty=tai2",
			input: input{
				tai1: models.Tai{},
				tai2: models.Tai{},
			},
			expect: false,
		},
		{
			name: "same: tai1=20893-000001=tai2",
			input: input{
				tai1: models.Tai{
					PlmnId: &models.PlmnId{
						Mcc: "208",
						Mnc: "93",
					},
					Tac: "000001",
				},
				tai2: models.Tai{
					PlmnId: &models.PlmnId{
						Mcc: "208",
						Mnc: "93",
					},
					Tac: "000001",
				},
			},
			expect: true,
		},
		{
			name: "diff: tai1=20893-000001, tai2=20894-000001",
			input: input{
				tai1: models.Tai{
					PlmnId: &models.PlmnId{
						Mcc: "208",
						Mnc: "93",
					},
					Tac: "000001",
				},
				tai2: models.Tai{
					PlmnId: &models.PlmnId{
						Mcc: "208",
						Mnc: "94",
					},
					Tac: "000001",
				},
			},
			expect: false,
		},
		{
			name: "diff: tai1=20893-000001, tai2=20893-000002",
			input: input{
				tai1: models.Tai{
					PlmnId: &models.PlmnId{
						Mcc: "208",
						Mnc: "93",
					},
					Tac: "000001",
				},
				tai2: models.Tai{
					PlmnId: &models.PlmnId{
						Mcc: "208",
						Mnc: "93",
					},
					Tac: "000002",
				},
			},
			expect: false,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.name, func(t *testing.T) {
			same := TaiEqualFold(tc.input.tai1, tc.input.tai2)
			require.Equal(t, tc.expect, same)
		})
	}
}

func TestGnbIdEqualFold(t *testing.T) {
	type input struct {
		gnbId1 models.GNbId
		gnbId2 models.GNbId
	}
	tcs := []struct {
		name   string
		input  input
		expect bool
	}{
		{
			name: "same: gnbId1=abcdef01/32=gnbId2",
			input: input{
				gnbId1: models.GNbId{
					BitLength: 32,
					GNBValue:  "abcdef01",
				},
				gnbId2: models.GNbId{
					BitLength: 32,
					GNBValue:  "abcdef01",
				},
			},
			expect: true,
		},
		{
			name: "same: gnbId1:abcdef01/32=gnbId2:ABCDEF01",
			input: input{
				gnbId1: models.GNbId{
					BitLength: 32,
					GNBValue:  "abcdef01",
				},
				gnbId2: models.GNbId{
					BitLength: 32,
					GNBValue:  "ABCDEF01",
				},
			},
			expect: true,
		},
		{
			name: "diff: gnbId1:abcdef01/31=gnbId2:abcdef01/32",
			input: input{
				gnbId1: models.GNbId{
					BitLength: 31,
					GNBValue:  "abcdef01",
				},
				gnbId2: models.GNbId{
					BitLength: 32,
					GNBValue:  "abcdef01",
				},
			},
			expect: false,
		},
		{
			name: "diff: gnbId1:abcdef01/32=gnbId2:abcdef02/32",
			input: input{
				gnbId1: models.GNbId{
					BitLength: 32,
					GNBValue:  "abcdef01",
				},
				gnbId2: models.GNbId{
					BitLength: 32,
					GNBValue:  "abcdef02",
				},
			},
			expect: false,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.name, func(t *testing.T) {
			result := GnbIdEqualFold(tc.input.gnbId1, tc.input.gnbId2)
			require.Equal(t, tc.expect, result)
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

func TestModelsGnbIdToUint(t *testing.T) {
	testCases := []struct {
		name     string
		gnbId    *models.GNbId
		expected uint64
	}{
		{
			name:     "nil input",
			gnbId:    nil,
			expected: 0,
		},
		{
			name: "pass1",
			gnbId: &models.GNbId{
				GNBValue:  "0000001a",
				BitLength: 32,
			},
			expected: 0x1a,
		},
		{
			name: "bitLen too short",
			gnbId: &models.GNbId{
				GNBValue:  "0001000a",
				BitLength: 16,
			},
			expected: 0,
		},
		{
			name: "invalid gnbValue",
			gnbId: &models.GNbId{
				GNBValue:  "0000000z",
				BitLength: 32,
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ModelsGnbIdToUint(tc.gnbId)
			if result == 0 {
				fmt.Println(err)
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestModelsPlmnIdNidToPlmnId(t *testing.T) {
	testCases := []struct {
		name      string
		plmnidNid *models.PlmnIdNid
		expected  *models.PlmnId
	}{
		{
			name:      "nil input",
			plmnidNid: nil,
			expected:  nil,
		},
		{
			name: "normal",
			plmnidNid: &models.PlmnIdNid{
				Mcc: "001",
				Mnc: "01",
				Nid: "1",
			},
			expected: &models.PlmnId{
				Mcc: "001",
				Mnc: "01",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ModelsPlmnIdNidToPlmnId(tc.plmnidNid)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestModelsPlmnIdToPlmnIdNid(t *testing.T) {
	testCases := []struct {
		name     string
		plmnId   *models.PlmnId
		expected *models.PlmnIdNid
	}{
		{
			name:     "nil input",
			plmnId:   nil,
			expected: nil,
		},
		{
			name: "normal",
			plmnId: &models.PlmnId{
				Mcc: "001",
				Mnc: "01",
			},
			expected: &models.PlmnIdNid{
				Mcc: "001",
				Mnc: "01",
				Nid: "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ModelsPlmnIdToPlmnIdNid(tc.plmnId)
			require.Equal(t, tc.expected, result)
		})
	}
}
