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

type ReportedEventType string

// List of ReportedEventType
const (
	ReportedEventType_PERIODIC_EVENT                    ReportedEventType = "PERIODIC_EVENT"
	ReportedEventType_ENTERING_AREA_EVENT               ReportedEventType = "ENTERING_AREA_EVENT"
	ReportedEventType_LEAVING_AREA_EVENT                ReportedEventType = "LEAVING_AREA_EVENT"
	ReportedEventType_BEING_INSIDE_AREA_EVENT           ReportedEventType = "BEING_INSIDE_AREA_EVENT"
	ReportedEventType_MOTION_EVENT                      ReportedEventType = "MOTION_EVENT"
	ReportedEventType_MAXIMUM_INTERVAL_EXPIRATION_EVENT ReportedEventType = "MAXIMUM_INTERVAL_EXPIRATION_EVENT"
	ReportedEventType_LOCATION_CANCELLATION_EVENT       ReportedEventType = "LOCATION_CANCELLATION_EVENT"
)
