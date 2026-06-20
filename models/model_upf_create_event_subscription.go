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

// UpfCreateEventSubscription is the request wrapper for creating an UPF event subscription.
type UpfCreateEventSubscription struct {
	Subscription      UpfEventSubscription `json:"subscription" yaml:"subscription" bson:"subscription,omitempty"`
	SupportedFeatures string               `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
}
