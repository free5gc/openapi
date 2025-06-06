/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.522 V17.7.0; 5G System; Network Exposure Function Northbound APIs.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.522/
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Indicates the time interval(s) during which the AF request is to be applied.
type TemporalValidity struct {
	// string with format 'date-time' as defined in OpenAPI.
	StartTime *time.Time `json:"startTime,omitempty" yaml:"startTime" bson:"startTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	StopTime *time.Time `json:"stopTime,omitempty" yaml:"stopTime" bson:"stopTime,omitempty"`
}
