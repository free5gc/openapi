/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Report Item for UE Location Trends event.
type UeLocationTrendsReportItem struct {
	Tai          *Tai          `json:"tai,omitempty" yaml:"tai" bson:"tai,omitempty"`
	Ncgi         *Ncgi         `json:"ncgi,omitempty" yaml:"ncgi" bson:"ncgi,omitempty"`
	Ecgi         *Ecgi         `json:"ecgi,omitempty" yaml:"ecgi" bson:"ecgi,omitempty"`
	N3gaLocation *N3gaLocation `json:"n3gaLocation,omitempty" yaml:"n3gaLocation" bson:"n3gaLocation,omitempty"`
	// indicating a time in seconds.
	Spacing int32 `json:"spacing" yaml:"spacing" bson:"spacing,omitempty"`
	// indicating a time in seconds.
	Duration int32 `json:"duration" yaml:"duration" bson:"duration,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	Timestamp *time.Time `json:"timestamp" yaml:"timestamp" bson:"timestamp,omitempty"`
}
