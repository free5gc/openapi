/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains context information related to analytics subscriptions corresponding with one or  more context identifiers.
type ContextData struct {
	// List of items that contain context information corresponding with a context identifier.
	ContextElems []ContextElement `json:"contextElems" yaml:"contextElems" bson:"contextElems,omitempty"`
}
