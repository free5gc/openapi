/*
 * LMF Location
 *
 * LMF Location Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.572 V17.9.0; 5G System; Location Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.572/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Indicates the information of periodic event reporting.
type PeriodicEventInfo struct {
	// Number of required periodic event reports.
	ReportingAmount int32 `json:"reportingAmount" yaml:"reportingAmount" bson:"reportingAmount,omitempty"`
	// Event reporting periodic interval in seconds.
	ReportingInterval    int32 `json:"reportingInterval" yaml:"reportingInterval" bson:"reportingInterval,omitempty"`
	ReportingInfiniteInd bool  `json:"reportingInfiniteInd,omitempty" yaml:"reportingInfiniteInd" bson:"reportingInfiniteInd,omitempty"`
	// Event reporting periodic interval in milliseconds.
	ReportingIntervalMs int32 `json:"reportingIntervalMs,omitempty" yaml:"reportingIntervalMs" bson:"reportingIntervalMs,omitempty"`
}
