/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the Spatial Validity Condition.
type SpatialValidityCond struct {
	TrackingAreaList        []Tai           `json:"trackingAreaList,omitempty" yaml:"trackingAreaList" bson:"trackingAreaList,omitempty"`
	Countries               []string        `json:"countries,omitempty" yaml:"countries" bson:"countries,omitempty"`
	GeographicalServiceArea *GeoServiceArea `json:"geographicalServiceArea,omitempty" yaml:"geographicalServiceArea" bson:"geographicalServiceArea,omitempty"`
}
