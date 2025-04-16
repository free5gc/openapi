/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Represents the expected UE behavior (e.g. UE moving trajectory) and its validity period
type ExpectedUeBehavior struct {
	ExpMoveTrajectory []UserLocation `json:"expMoveTrajectory" yaml:"expMoveTrajectory" bson:"expMoveTrajectory,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime" yaml:"validityTime" bson:"validityTime,omitempty"`
}
