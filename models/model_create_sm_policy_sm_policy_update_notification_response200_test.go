package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateSMPolicySmPolicyUpdateNotificationResponse200JSON(t *testing.T) {
	partialSuccess := `[{"failureCause":"PCC_RULE_EVENT","ruleReports":[{"pccRuleIds":["PccRule-2"],` +
		`"ruleStatus":"INACTIVE","failureCode":"RES_ALLO_FAIL"}]}]`
	ueCampingRep := `{"ueTimeZone":"+08:00","netLocAccSupp":"LOC_NOT_SUPPORTED"}`

	tests := []struct {
		name  string
		input string
		want  CreateSMPolicySmPolicyUpdateNotificationResponse200
	}{
		{
			name:  "UeCampingRep",
			input: ueCampingRep,
			want: CreateSMPolicySmPolicyUpdateNotificationResponse200{
				UeTimeZone:    "+08:00",
				NetLocAccSupp: Pcf_SMPolCtrl_NetLocAccessSupport_LOC_NOT_SUPPORTED,
			},
		},
		{
			name:  "array of PartialSuccessReport",
			input: partialSuccess,
			want: CreateSMPolicySmPolicyUpdateNotificationResponse200{
				PartialSuccessReports: []Pcf_SMPolCtrl_PartialSuccessReport{{
					FailureCause: Pcf_SMPolCtrl_FailureCause_PCC_RULE_EVENT,
					RuleReports: []Pcf_SMPolCtrl_RuleReport{{
						PccRuleIds:  []string{"PccRule-2"},
						RuleStatus:  Pcf_SMPolCtrl_RuleStatus_INACTIVE,
						FailureCode: Pcf_SMPolCtrl_FailureCode_RES_ALLO_FAIL,
					}},
				}},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got CreateSMPolicySmPolicyUpdateNotificationResponse200
			require.NoError(t, json.Unmarshal([]byte(tc.input), &got))
			require.Equal(t, tc.want, got)

			encoded, err := json.Marshal(got)
			require.NoError(t, err)
			require.JSONEq(t, tc.input, string(encoded))
		})
	}
}

func TestCreateSMPolicySmPolicyUpdateNotificationResponse200InvalidJSON(t *testing.T) {
	for _, input := range []string{`[]`, `"UeCampingRep"`, `[{"ruleReports":{}}]`} {
		var got CreateSMPolicySmPolicyUpdateNotificationResponse200
		require.Error(t, json.Unmarshal([]byte(input), &got), input)
	}
}
