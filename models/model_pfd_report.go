/*
 * 3gpp-pfd-management
 *
 * API for PFD management.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * Source file: 3GPP TS 29.122 V17.10.0 T8 reference point for Northbound APIs
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.122/
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	
)



// Represents a PFD report indicating the external application identifier(s) which PFD(s) are not added or modified successfully and the corresponding failure cause(s).
type PfdReport struct {
	// Identifies the external application identifier(s) which PFD(s) are not added or modified successfully
	ExternalAppIds []string `json:"externalAppIds,omitempty" yaml:"externalAppIds" bson:"externalAppIds,omitempty"`
	FailureCode FailureCode `json:"failureCode,omitempty" yaml:"failureCode" bson:"failureCode,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	CachingTime int32 `json:"cachingTime,omitempty" yaml:"cachingTime" bson:"cachingTime,omitempty"`
	LocationArea *UserPlaneLocationArea `json:"locationArea,omitempty" yaml:"locationArea" bson:"locationArea,omitempty"`
}

