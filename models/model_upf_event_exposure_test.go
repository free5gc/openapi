package models

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSmfUpfEventSubscriptionWireShape(t *testing.T) {
	bundleId := uint32(0)
	subscription := SmfEventExposureEventSubscription{
		Event: SmfEvent_UPF_EVENT,
		UpfEvents: []UpfEvent{
			{
				Type:                     UpfEventType_USER_DATA_USAGE_MEASURES,
				MeasurementTypes:         []UpfMeasurementType{UpfMeasurementType_VOLUME_MEASUREMENT},
				GranularityOfMeasurement: UpfGranularityOfMeasurement_PER_SESSION,
			},
		},
		BundlingAllowed:       true,
		BundleId:              &bundleId,
		BundledEventNotifyUri: "https://nwdaf.example.com/upf-events",
	}

	got, err := json.Marshal(subscription)
	require.NoError(t, err)
	require.JSONEq(t, `{
		"event":"UPF_EVENT",
		"upfEvents":[{
			"type":"USER_DATA_USAGE_MEASURES",
			"measurementTypes":["VOLUME_MEASUREMENT"],
			"granularityOfMeasurement":"PER_SESSION"
		}],
		"bundlingAllowed":true,
		"bundleId":0,
		"bundledEventNotifyUri":"https://nwdaf.example.com/upf-events"
	}`, string(got))
}

func TestUpfCreateEventSubscriptionWireShape(t *testing.T) {
	bundleId := uint32(0)
	request := UpfCreateEventSubscription{
		Subscription: UpfEventSubscription{
			EventList: []UpfEvent{
				{
					Type: UpfEventType_USER_DATA_USAGE_MEASURES,
					MeasurementTypes: []UpfMeasurementType{
						UpfMeasurementType_VOLUME_MEASUREMENT,
						UpfMeasurementType_THROUGHPUT_MEASUREMENT,
					},
					GranularityOfMeasurement: UpfGranularityOfMeasurement_PER_SESSION,
				},
			},
			EventNotifyUri:      "https://nwdaf.example.com/upf-events",
			NotifyCorrelationId: "nsmf-notif-id",
			EventReportingMode: UpfEventMode{
				Trigger:   UpfEventTrigger_PERIODIC,
				RepPeriod: 10,
			},
			NfId:                  "11111111-2222-3333-4444-555555555555",
			UeIpAddress:           &IpAddr{Ipv4Addr: "192.0.2.1"},
			Supi:                  "imsi-001010000000001",
			Dnn:                   "internet",
			Snssai:                &Snssai{Sst: 1, Sd: "010203"},
			BundlingAllowed:       true,
			BundleId:              &bundleId,
			BundledEventNotifyUri: "https://nwdaf.example.com/bundled-upf-events",
		},
		SupportedFeatures: "1",
	}

	got, err := json.Marshal(request)
	require.NoError(t, err)
	require.JSONEq(t, `{
		"subscription":{
			"eventList":[{
				"type":"USER_DATA_USAGE_MEASURES",
				"measurementTypes":["VOLUME_MEASUREMENT","THROUGHPUT_MEASUREMENT"],
				"granularityOfMeasurement":"PER_SESSION"
			}],
			"eventNotifyUri":"https://nwdaf.example.com/upf-events",
			"notifyCorrelationId":"nsmf-notif-id",
			"eventReportingMode":{"trigger":"PERIODIC","repPeriod":10},
			"nfId":"11111111-2222-3333-4444-555555555555",
			"ueIpAddress":{"ipv4Addr":"192.0.2.1"},
			"supi":"imsi-001010000000001",
			"dnn":"internet",
			"snssai":{"sst":1,"sd":"010203"},
			"bundlingAllowed":true,
			"bundleId":0,
			"bundledEventNotifyUri":"https://nwdaf.example.com/bundled-upf-events"
		},
		"supportedFeatures":"1"
	}`, string(got))
}

func TestUpfEventSubscriptionAnyUeWireShape(t *testing.T) {
	subscription := UpfEventSubscription{
		EventList:           []UpfEvent{{Type: UpfEventType_USER_DATA_USAGE_MEASURES}},
		EventNotifyUri:      "https://nwdaf.example.com/upf-events",
		NotifyCorrelationId: "nsmf-notif-id",
		EventReportingMode:  UpfEventMode{Trigger: UpfEventTrigger_PERIODIC},
		NfId:                "11111111-2222-3333-4444-555555555555",
		AnyUe:               true,
	}

	got, err := json.Marshal(subscription)
	require.NoError(t, err)
	require.Contains(t, string(got), `"anyUe":true`)
}

func TestBundleIdUint32UpperBoundWireShape(t *testing.T) {
	bundleId := uint32(4294967295)
	tests := []struct {
		name  string
		value interface{}
	}{
		{
			name: "Nsmf",
			value: SmfEventExposureEventSubscription{
				Event:    SmfEvent_UPF_EVENT,
				BundleId: &bundleId,
			},
		},
		{
			name: "Nupf",
			value: UpfEventSubscription{
				EventList:           []UpfEvent{{Type: UpfEventType_USER_DATA_USAGE_MEASURES}},
				EventNotifyUri:      "https://nwdaf.example.com/upf-events",
				NotifyCorrelationId: "nsmf-notif-id",
				EventReportingMode:  UpfEventMode{Trigger: UpfEventTrigger_PERIODIC},
				NfId:                "11111111-2222-3333-4444-555555555555",
				BundleId:            &bundleId,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := json.Marshal(test.value)
			require.NoError(t, err)
			require.Contains(t, string(got), `"bundleId":4294967295`)
		})
	}
}

func TestBundleIdRejectsValuesOutsideUint32Range(t *testing.T) {
	tests := []struct {
		name      string
		payload   string
		newTarget func() interface{}
	}{
		{
			name:      "Nsmf negative",
			payload:   `{"event":"UPF_EVENT","bundleId":-1}`,
			newTarget: func() interface{} { return &SmfEventExposureEventSubscription{} },
		},
		{
			name:      "Nsmf above maximum",
			payload:   `{"event":"UPF_EVENT","bundleId":4294967296}`,
			newTarget: func() interface{} { return &SmfEventExposureEventSubscription{} },
		},
		{
			name:      "Nupf negative",
			payload:   `{"bundleId":-1}`,
			newTarget: func() interface{} { return &UpfEventSubscription{} },
		},
		{
			name:      "Nupf above maximum",
			payload:   `{"bundleId":4294967296}`,
			newTarget: func() interface{} { return &UpfEventSubscription{} },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := json.Unmarshal([]byte(test.payload), test.newTarget())
			require.Error(t, err)
		})
	}
}

func TestBundlingAllowedFalseIsOmitted(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{
			name:  "Nsmf",
			value: SmfEventExposureEventSubscription{Event: SmfEvent_UPF_EVENT},
		},
		{
			name: "Nupf",
			value: UpfEventSubscription{
				EventList:          []UpfEvent{{Type: UpfEventType_USER_DATA_USAGE_MEASURES}},
				EventNotifyUri:     "https://nwdaf.example.com/upf-events",
				EventReportingMode: UpfEventMode{Trigger: UpfEventTrigger_PERIODIC},
				NfId:               "11111111-2222-3333-4444-555555555555",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := json.Marshal(test.value)
			require.NoError(t, err)
			require.NotContains(t, string(got), "bundlingAllowed")
		})
	}
}

func TestUpfEventExposureEnumValues(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{"event QOS_MONITORING", UpfEventType_QOS_MONITORING, "QOS_MONITORING"},
		{"event USER_DATA_USAGE_MEASURES", UpfEventType_USER_DATA_USAGE_MEASURES, "USER_DATA_USAGE_MEASURES"},
		{"event USER_DATA_USAGE_TRENDS", UpfEventType_USER_DATA_USAGE_TRENDS, "USER_DATA_USAGE_TRENDS"},
		{"event TSC_MNGT_INFO", UpfEventType_TSC_MNGT_INFO, "TSC_MNGT_INFO"},
		{"event UE_NAT_MAPPING_INFO", UpfEventType_UE_NAT_MAPPING_INFO, "UE_NAT_MAPPING_INFO"},
		{"event HANDLING_OF_PAYLOAD_HEADERS_INFO", UpfEventType_HANDLING_OF_PAYLOAD_HEADERS_INFO, "HANDLING_OF_PAYLOAD_HEADERS_INFO"},
		{"event SUBSCRIPTION_TERMINATION", UpfEventType_SUBSCRIPTION_TERMINATION, "SUBSCRIPTION_TERMINATION"},
		{"measurement VOLUME_MEASUREMENT", UpfMeasurementType_VOLUME_MEASUREMENT, "VOLUME_MEASUREMENT"},
		{"measurement THROUGHPUT_MEASUREMENT", UpfMeasurementType_THROUGHPUT_MEASUREMENT, "THROUGHPUT_MEASUREMENT"},
		{"measurement APPLICATION_RELATED_INFO", UpfMeasurementType_APPLICATION_RELATED_INFO, "APPLICATION_RELATED_INFO"},
		{"granularity PER_APPLICATION", UpfGranularityOfMeasurement_PER_APPLICATION, "PER_APPLICATION"},
		{"granularity PER_SESSION", UpfGranularityOfMeasurement_PER_SESSION, "PER_SESSION"},
		{"granularity PER_FLOW", UpfGranularityOfMeasurement_PER_FLOW, "PER_FLOW"},
		{"trigger ONE_TIME", UpfEventTrigger_ONE_TIME, "ONE_TIME"},
		{"trigger PERIODIC", UpfEventTrigger_PERIODIC, "PERIODIC"},
		{"trigger CONTINUOUS", UpfEventTrigger_CONTINUOUS, "CONTINUOUS"},
		{"service NUPF_EE", ServiceName_NUPF_EE, "nupf-ee"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.want, fmt.Sprint(test.value))
		})
	}
}
