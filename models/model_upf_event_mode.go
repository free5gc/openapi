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

// UpfEventMode represents the reporting controls used by the SMF Event Exposure profile.
type UpfEventMode struct {
	Trigger   UpfEventTrigger `json:"trigger" yaml:"trigger" bson:"trigger,omitempty"`
	RepPeriod int32           `json:"repPeriod,omitempty" yaml:"repPeriod" bson:"repPeriod,omitempty"`
}
