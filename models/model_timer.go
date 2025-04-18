/*
 * Nudsf_Timer
 *
 * Nudsf Timer Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.598 UDSF Services, V17.6.0.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.598/
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Represents a timer.
type Timer struct {
	// Represents the identifier of a timer.
	TimerId string `json:"timerId,omitempty" yaml:"timerId" bson:"timerId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expires *time.Time `json:"expires" yaml:"expires" bson:"expires,omitempty"`
	// A map (list of key-value pairs where a tagName of type string serves as key) of tagValue lists
	MetaTags map[string][]string `json:"metaTags,omitempty" yaml:"metaTags" bson:"metaTags,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	CallbackReference string `json:"callbackReference,omitempty" yaml:"callbackReference" bson:"callbackReference,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	DeleteAfter int32 `json:"deleteAfter,omitempty" yaml:"deleteAfter" bson:"deleteAfter,omitempty"`
}
