/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.7.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents the S-NSSAI and the optionally associated Network Slice Instance(s).
type NsiIdInfo struct {
	Snssai *Snssai  `json:"snssai" yaml:"snssai" bson:"snssai,omitempty"`
	NsiIds []string `json:"nsiIds,omitempty" yaml:"nsiIds" bson:"nsiIds,omitempty"`
}
