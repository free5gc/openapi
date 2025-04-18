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

// Represents UE communication information.
type UeCommunication struct {
	// indicating a time in seconds.
	CommDur int32 `json:"commDur,omitempty" yaml:"commDur" bson:"commDur,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	CommDurVariance float32 `json:"commDurVariance,omitempty" yaml:"commDurVariance" bson:"commDurVariance,omitempty"`
	// indicating a time in seconds.
	PerioTime int32 `json:"perioTime,omitempty" yaml:"perioTime" bson:"perioTime,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	PerioTimeVariance float32 `json:"perioTimeVariance,omitempty" yaml:"perioTimeVariance" bson:"perioTimeVariance,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Ts *time.Time `json:"ts,omitempty" yaml:"ts" bson:"ts,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	TsVariance    float32                     `json:"tsVariance,omitempty" yaml:"tsVariance" bson:"tsVariance,omitempty"`
	RecurringTime *ScheduledCommunicationTime `json:"recurringTime,omitempty" yaml:"recurringTime" bson:"recurringTime,omitempty"`
	TrafChar      *TrafficCharacterization    `json:"trafChar,omitempty" yaml:"trafChar" bson:"trafChar,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	Ratio        int32 `json:"ratio,omitempty" yaml:"ratio" bson:"ratio,omitempty"`
	PerioCommInd bool  `json:"perioCommInd,omitempty" yaml:"perioCommInd" bson:"perioCommInd,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence     int32                    `json:"confidence,omitempty" yaml:"confidence" bson:"confidence,omitempty"`
	AnaOfAppList   *AppListForUeComm        `json:"anaOfAppList,omitempty" yaml:"anaOfAppList" bson:"anaOfAppList,omitempty"`
	SessInactTimer *SessInactTimerForUeComm `json:"sessInactTimer,omitempty" yaml:"sessInactTimer" bson:"sessInactTimer,omitempty"`
}
