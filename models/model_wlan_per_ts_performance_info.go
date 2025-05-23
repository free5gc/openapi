/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.10.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// WLAN performance information per Time Slot during the analytics target period.
type WlanPerTsPerformanceInfo struct {
	// string with format 'date-time' as defined in OpenAPI.
	TsStart *time.Time `json:"tsStart" yaml:"tsStart" bson:"tsStart,omitempty"`
	// indicating a time in seconds.
	TsDuration int32 `json:"tsDuration" yaml:"tsDuration" bson:"tsDuration,omitempty"`
	Rssi       int32 `json:"rssi,omitempty" yaml:"rssi" bson:"rssi,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Rtt         int32               `json:"rtt,omitempty" yaml:"rtt" bson:"rtt,omitempty"`
	TrafficInfo *TrafficInformation `json:"trafficInfo,omitempty" yaml:"trafficInfo" bson:"trafficInfo,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	NumberOfUes int32 `json:"numberOfUes,omitempty" yaml:"numberOfUes" bson:"numberOfUes,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty" yaml:"confidence" bson:"confidence,omitempty"`
}
