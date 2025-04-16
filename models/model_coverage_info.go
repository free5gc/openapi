/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291 V17.9.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 3.1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type CoverageInfo struct {
	CoverageStatus bool `json:"coverageStatus,omitempty" yaml:"coverageStatus" bson:"coverageStatus,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ChangeTime   *time.Time     `json:"changeTime,omitempty" yaml:"changeTime" bson:"changeTime,omitempty"`
	LocationInfo []UserLocation `json:"locationInfo,omitempty" yaml:"locationInfo" bson:"locationInfo,omitempty"`
}
