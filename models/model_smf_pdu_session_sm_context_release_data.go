/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V17.11.0; 5G System; Session Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Data within Release SM Context Request
type SmfPduSessionSmContextReleaseData struct {
	Cause     SmfPduSessionCause `json:"cause,omitempty" yaml:"cause" bson:"cause,omitempty"`
	NgApCause *NgApCause         `json:"ngApCause,omitempty" yaml:"ngApCause" bson:"ngApCause,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var5gMmCauseValue int32         `json:"5gMmCauseValue,omitempty" yaml:"5gMmCauseValue" bson:"5gMmCauseValue,omitempty"`
	UeLocation        *UserLocation `json:"ueLocation,omitempty" yaml:"ueLocation" bson:"ueLocation,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	UeTimeZone      string           `json:"ueTimeZone,omitempty" yaml:"ueTimeZone" bson:"ueTimeZone,omitempty"`
	AddUeLocation   *UserLocation    `json:"addUeLocation,omitempty" yaml:"addUeLocation" bson:"addUeLocation,omitempty"`
	VsmfReleaseOnly bool             `json:"vsmfReleaseOnly,omitempty" yaml:"vsmfReleaseOnly" bson:"vsmfReleaseOnly,omitempty"`
	N2SmInfo        *RefToBinaryData `json:"n2SmInfo,omitempty" yaml:"n2SmInfo" bson:"n2SmInfo,omitempty"`
	N2SmInfoType    N2SmInfoType     `json:"n2SmInfoType,omitempty" yaml:"n2SmInfoType" bson:"n2SmInfoType,omitempty"`
	IsmfReleaseOnly bool             `json:"ismfReleaseOnly,omitempty" yaml:"ismfReleaseOnly" bson:"ismfReleaseOnly,omitempty"`
}
