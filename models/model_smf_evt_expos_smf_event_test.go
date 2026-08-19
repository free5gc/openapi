package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSmfEvtExposSmfEventUPFEventWireValue(t *testing.T) {
	require.Equal(t, Smf_EvtExpos_SmfEvent("UPF_EVENT"), Smf_EvtExpos_SmfEvent_UPF_EVENT)

	wireValue, err := json.Marshal(Smf_EvtExpos_SmfEvent_UPF_EVENT)
	require.NoError(t, err)
	require.Equal(t, `"UPF_EVENT"`, string(wireValue))
}
