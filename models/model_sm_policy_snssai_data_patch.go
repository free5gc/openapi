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

// Contains the SM policy data for a given subscriber and S-NSSAI.
type SmPolicySnssaiDataPatch struct {
	Snssai *Snssai `json:"snssai" yaml:"snssai" bson:"snssai,omitempty"`
	// Modifiable Session Management Policy data per DNN for all the DNNs of the indicated S-NSSAI. The key of the map is the DNN.
	SmPolicyDnnData map[string]SmPolicyDnnDataPatch `json:"smPolicyDnnData,omitempty" yaml:"smPolicyDnnData" bson:"smPolicyDnnData,omitempty"`
}
