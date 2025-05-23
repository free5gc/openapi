/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.514 V17.9.0; 5G System; Policy Authorization Service; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.514/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Indicates the start or stop of the detected application traffic and the application identifier of the detected application traffic.
type AppDetectionReport struct {
	AdNotifType AppDetectionNotifType `json:"adNotifType" yaml:"adNotifType" bson:"adNotifType,omitempty"`
	// Contains an AF application identifier.
	AfAppId string `json:"afAppId" yaml:"afAppId" bson:"afAppId,omitempty"`
}
