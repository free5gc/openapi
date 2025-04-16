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

// This data type is defined in the same way as the QosMonitoringInformation data type, but with the OpenAPI nullable property set to true.
type PcfPolicyAuthorizationQosMonitoringInformationRm struct {
	RepThreshDl int32 `json:"repThreshDl,omitempty" yaml:"repThreshDl" bson:"repThreshDl,omitempty"`
	RepThreshUl int32 `json:"repThreshUl,omitempty" yaml:"repThreshUl" bson:"repThreshUl,omitempty"`
	RepThreshRp int32 `json:"repThreshRp,omitempty" yaml:"repThreshRp" bson:"repThreshRp,omitempty"`
}
