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

// UpfEventSubscription represents an UPF Event Exposure subscription.
type UpfEventSubscription struct {
	EventList             []UpfEvent   `json:"eventList" yaml:"eventList" bson:"eventList,omitempty"`
	EventNotifyUri        string       `json:"eventNotifyUri" yaml:"eventNotifyUri" bson:"eventNotifyUri,omitempty"`
	NotifyCorrelationId   string       `json:"notifyCorrelationId" yaml:"notifyCorrelationId" bson:"notifyCorrelationId,omitempty"`
	EventReportingMode    UpfEventMode `json:"eventReportingMode" yaml:"eventReportingMode" bson:"eventReportingMode,omitempty"`
	NfId                  string       `json:"nfId" yaml:"nfId" bson:"nfId,omitempty"`
	UeIpAddress           *IpAddr      `json:"ueIpAddress,omitempty" yaml:"ueIpAddress" bson:"ueIpAddress,omitempty"`
	AnyUe                 bool         `json:"anyUe,omitempty" yaml:"anyUe" bson:"anyUe,omitempty"`
	Supi                  string       `json:"supi,omitempty" yaml:"supi" bson:"supi,omitempty"`
	Dnn                   string       `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	Snssai                *Snssai      `json:"snssai,omitempty" yaml:"snssai" bson:"snssai,omitempty"`
	BundlingAllowed       bool         `json:"bundlingAllowed,omitempty" yaml:"bundlingAllowed" bson:"bundlingAllowed,omitempty"`
	BundleId              *uint32      `json:"bundleId,omitempty" yaml:"bundleId" bson:"bundleId,omitempty"`
	BundledEventNotifyUri string       `json:"bundledEventNotifyUri,omitempty" yaml:"bundledEventNotifyUri" bson:"bundledEventNotifyUri,omitempty"`
}
