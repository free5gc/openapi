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

type UpfGranularityOfMeasurement string

// List of UpfGranularityOfMeasurement
const (
	UpfGranularityOfMeasurement_PER_APPLICATION UpfGranularityOfMeasurement = "PER_APPLICATION"
	UpfGranularityOfMeasurement_PER_SESSION     UpfGranularityOfMeasurement = "PER_SESSION"
	UpfGranularityOfMeasurement_PER_FLOW        UpfGranularityOfMeasurement = "PER_FLOW"
)
