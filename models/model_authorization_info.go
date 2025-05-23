/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Represents NIDD authorization information.
type AuthorizationInfo struct {
	Snssai *Snssai `json:"snssai" yaml:"snssai" bson:"snssai,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn" yaml:"dnn" bson:"dnn,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderInformation string `json:"mtcProviderInformation" yaml:"mtcProviderInformation" bson:"mtcProviderInformation,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	AuthUpdateCallbackUri string `json:"authUpdateCallbackUri" yaml:"authUpdateCallbackUri" bson:"authUpdateCallbackUri,omitempty"`
	AfId                  string `json:"afId,omitempty" yaml:"afId" bson:"afId,omitempty"`
	// Identity of the NEF
	NefId string `json:"nefId,omitempty" yaml:"nefId" bson:"nefId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time   `json:"validityTime,omitempty" yaml:"validityTime" bson:"validityTime,omitempty"`
	ContextInfo  *ContextInfo `json:"contextInfo,omitempty" yaml:"contextInfo" bson:"contextInfo,omitempty"`
}
