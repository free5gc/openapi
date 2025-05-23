/*
 * AUSF API
 *
 * AUSF UE Authentication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.509 V17.9.0; 5G System; 3GPP TS Authentication Server services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.509
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains Authentication Vector for method 5G AKA.
type Av5gAka struct {
	Rand string `json:"rand" yaml:"rand" bson:"rand,omitempty"`
	// Contains the HXRES*.
	HxresStar string `json:"hxresStar" yaml:"hxresStar" bson:"hxresStar,omitempty"`
	Autn      string `json:"autn" yaml:"autn" bson:"autn,omitempty"`
}
