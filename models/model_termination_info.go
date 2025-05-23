/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.514 V17.9.0; 5G System; Policy Authorization Service; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.514/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Indicates the cause for requesting the deletion of the Individual Application Session Context resource.
type TerminationInfo struct {
	TermCause PcfPolicyAuthorizationTerminationCause `json:"termCause" yaml:"termCause" bson:"termCause,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	ResUri string `json:"resUri" yaml:"resUri" bson:"resUri,omitempty"`
}
