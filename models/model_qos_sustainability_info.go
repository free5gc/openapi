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

// Represents the QoS Sustainability information.
type QosSustainabilityInfo struct {
	AreaInfo *NetworkAreaInfo `json:"areaInfo,omitempty" yaml:"areaInfo" bson:"areaInfo,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	StartTs *time.Time `json:"startTs,omitempty" yaml:"startTs" bson:"startTs,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	EndTs         *time.Time              `json:"endTs,omitempty" yaml:"endTs" bson:"endTs,omitempty"`
	QosFlowRetThd *RetainabilityThreshold `json:"qosFlowRetThd,omitempty" yaml:"qosFlowRetThd" bson:"qosFlowRetThd,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	RanUeThrouThd string  `json:"ranUeThrouThd,omitempty" yaml:"ranUeThrouThd" bson:"ranUeThrouThd,omitempty"`
	Snssai        *Snssai `json:"snssai,omitempty" yaml:"snssai" bson:"snssai,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty" yaml:"confidence" bson:"confidence,omitempty"`
}
