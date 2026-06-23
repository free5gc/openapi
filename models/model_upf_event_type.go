/*
 * Nupf_EventExposure
 *
 * UPF Event Exposure Service.   Copyright 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.564 V19.5.0; 5G System; User Plane Function Services; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.564/
 *
 * API version: 1.2.0
 */

package models

type UpfEventType string

// List of UpfEventType
const (
	UpfEventType_QOS_MONITORING                   UpfEventType = "QOS_MONITORING"
	UpfEventType_USER_DATA_USAGE_MEASURES         UpfEventType = "USER_DATA_USAGE_MEASURES"
	UpfEventType_USER_DATA_USAGE_TRENDS           UpfEventType = "USER_DATA_USAGE_TRENDS"
	UpfEventType_TSC_MNGT_INFO                    UpfEventType = "TSC_MNGT_INFO"
	UpfEventType_UE_NAT_MAPPING_INFO              UpfEventType = "UE_NAT_MAPPING_INFO"
	UpfEventType_HANDLING_OF_PAYLOAD_HEADERS_INFO UpfEventType = "HANDLING_OF_PAYLOAD_HEADERS_INFO"
	UpfEventType_SUBSCRIPTION_TERMINATION         UpfEventType = "SUBSCRIPTION_TERMINATION"
)
