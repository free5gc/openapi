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

// Indicates the usage ranking criterion between the high, medium and low usage UE.
type RankingCriterion struct {
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	HighBase int32 `json:"highBase" yaml:"highBase" bson:"highBase,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	LowBase int32 `json:"lowBase" yaml:"lowBase" bson:"lowBase,omitempty"`
}
