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

type EdgeInfrastructureUsageChargingInformation struct {
	// string with format 'float' as defined in OpenAPI.
	MeanVirtualCPUUsage float32 `json:"meanVirtualCPUUsage,omitempty" yaml:"meanVirtualCPUUsage" bson:"meanVirtualCPUUsage,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	MeanVirtualMemoryUsage float32 `json:"meanVirtualMemoryUsage,omitempty" yaml:"meanVirtualMemoryUsage" bson:"meanVirtualMemoryUsage,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	MeanVirtualDiskUsage float32 `json:"meanVirtualDiskUsage,omitempty" yaml:"meanVirtualDiskUsage" bson:"meanVirtualDiskUsage,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.
	MeasuredInBytes int32 `json:"measuredInBytes,omitempty" yaml:"measuredInBytes" bson:"measuredInBytes,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.
	MeasuredOutBytes int32 `json:"measuredOutBytes,omitempty" yaml:"measuredOutBytes" bson:"measuredOutBytes,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	DurationStartTime *time.Time `json:"durationStartTime,omitempty" yaml:"durationStartTime" bson:"durationStartTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	DurationEndTime *time.Time `json:"durationEndTime,omitempty" yaml:"durationEndTime" bson:"durationEndTime,omitempty"`
}
