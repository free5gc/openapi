/*
 * 3gpp-ue-id
 *
 * API for UE Identifier Translation.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.522 V17.10.0 5G System; Network Exposure Function Northbound APIs
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.522/
 *
 * API version: 1.0.0
 */

package models

// UE identifier translation response data.
type Nef_UEId_UeIdTranslationRspData struct {
	// Subscription Permanent Identifier of the UE.
	Supi string `json:"supi,omitempty" yaml:"supi,omitempty" bson:"supi,omitempty"`

	// External identifier or MSISDN of the UE (GPSI).
	Gpsi string `json:"gpsi,omitempty" yaml:"gpsi,omitempty" bson:"gpsi,omitempty"`

	// Permanent Equipment Identifier of the UE.
	Pei string `json:"pei,omitempty" yaml:"pei,omitempty" bson:"pei,omitempty"`
}
