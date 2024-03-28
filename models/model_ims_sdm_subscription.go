/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.562 HSS Services for interworking with IMS, version 17.6.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.562/
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// A subscription to notifications of data change
type ImsSdmSubscription struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NfInstanceId string `json:"nfInstanceId" yaml:"nfInstanceId" bson:"nfInstanceId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	CallbackReference     string   `json:"callbackReference" yaml:"callbackReference" bson:"callbackReference,omitempty"`
	MonitoredResourceUris []string `json:"monitoredResourceUris" yaml:"monitoredResourceUris" bson:"monitoredResourceUris,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expires *time.Time `json:"expires,omitempty" yaml:"expires" bson:"expires,omitempty"`
}