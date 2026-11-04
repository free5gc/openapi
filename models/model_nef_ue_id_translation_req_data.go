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

// UE identifier translation request data.
type Nef_UEId_UeIdTranslationReqData struct {
	// External identifier or MSISDN of the UE (GPSI).
	Gpsi string `json:"gpsi,omitempty" yaml:"gpsi,omitempty" bson:"gpsi,omitempty"`

	// Subscription Concealed Identifier of the UE.
	Suci string `json:"suci,omitempty" yaml:"suci,omitempty" bson:"suci,omitempty"`
}
