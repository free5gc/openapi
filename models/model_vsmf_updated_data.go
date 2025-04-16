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

// Data within Update Response from V-SMF, or from I-SMF to SMF
type VsmfUpdatedData struct {
	QosFlowsAddModList         []QosFlowItem    `json:"qosFlowsAddModList,omitempty" yaml:"qosFlowsAddModList" bson:"qosFlowsAddModList,omitempty"`
	QosFlowsRelList            []QosFlowItem    `json:"qosFlowsRelList,omitempty" yaml:"qosFlowsRelList" bson:"qosFlowsRelList,omitempty"`
	QosFlowsFailedtoAddModList []QosFlowItem    `json:"qosFlowsFailedtoAddModList,omitempty" yaml:"qosFlowsFailedtoAddModList" bson:"qosFlowsFailedtoAddModList,omitempty"`
	QosFlowsFailedtoRelList    []QosFlowItem    `json:"qosFlowsFailedtoRelList,omitempty" yaml:"qosFlowsFailedtoRelList" bson:"qosFlowsFailedtoRelList,omitempty"`
	N1SmInfoFromUe             *RefToBinaryData `json:"n1SmInfoFromUe,omitempty" yaml:"n1SmInfoFromUe" bson:"n1SmInfoFromUe,omitempty"`
	UnknownN1SmInfo            *RefToBinaryData `json:"unknownN1SmInfo,omitempty" yaml:"unknownN1SmInfo" bson:"unknownN1SmInfo,omitempty"`
	UeLocation                 *UserLocation    `json:"ueLocation,omitempty" yaml:"ueLocation" bson:"ueLocation,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	UeTimeZone                  string                    `json:"ueTimeZone,omitempty" yaml:"ueTimeZone" bson:"ueTimeZone,omitempty"`
	AddUeLocation               *UserLocation             `json:"addUeLocation,omitempty" yaml:"addUeLocation" bson:"addUeLocation,omitempty"`
	AssignedEbiList             []EbiArpMapping           `json:"assignedEbiList,omitempty" yaml:"assignedEbiList" bson:"assignedEbiList,omitempty"`
	FailedToAssignEbiList       []Arp                     `json:"failedToAssignEbiList,omitempty" yaml:"failedToAssignEbiList" bson:"failedToAssignEbiList,omitempty"`
	ReleasedEbiList             []int32                   `json:"releasedEbiList,omitempty" yaml:"releasedEbiList" bson:"releasedEbiList,omitempty"`
	SecondaryRatUsageReport     []SecondaryRatUsageReport `json:"secondaryRatUsageReport,omitempty" yaml:"secondaryRatUsageReport" bson:"secondaryRatUsageReport,omitempty"`
	SecondaryRatUsageInfo       []SecondaryRatUsageInfo   `json:"secondaryRatUsageInfo,omitempty" yaml:"secondaryRatUsageInfo" bson:"secondaryRatUsageInfo,omitempty"`
	N4Info                      *N4Information            `json:"n4Info,omitempty" yaml:"n4Info" bson:"n4Info,omitempty"`
	N4InfoExt1                  *N4Information            `json:"n4InfoExt1,omitempty" yaml:"n4InfoExt1" bson:"n4InfoExt1,omitempty"`
	N4InfoExt2                  *N4Information            `json:"n4InfoExt2,omitempty" yaml:"n4InfoExt2" bson:"n4InfoExt2,omitempty"`
	N4InfoExt3                  *N4Information            `json:"n4InfoExt3,omitempty" yaml:"n4InfoExt3" bson:"n4InfoExt3,omitempty"`
	ModifiedEbiListNotDelivered bool                      `json:"modifiedEbiListNotDelivered,omitempty" yaml:"modifiedEbiListNotDelivered" bson:"modifiedEbiListNotDelivered,omitempty"`
}
