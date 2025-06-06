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

// Local 3D point with uncertainty ellipsoid
type Local3dPointUncertaintyEllipsoid struct {
	Shape                SupportedGadShapes         `json:"shape" yaml:"shape" bson:"shape,omitempty"`
	LocalOrigin          *LocalOrigin               `json:"localOrigin" yaml:"localOrigin" bson:"localOrigin,omitempty"`
	Point                *RelativeCartesianLocation `json:"point" yaml:"point" bson:"point,omitempty"`
	UncertaintyEllipsoid *UncertaintyEllipsoid      `json:"uncertaintyEllipsoid" yaml:"uncertaintyEllipsoid" bson:"uncertaintyEllipsoid,omitempty"`
	// Indicates value of confidence.
	Confidence int32 `json:"confidence" yaml:"confidence" bson:"confidence,omitempty"`
}
