/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Information of an AMF NF Instance
type NrfNfManagementAmfInfo struct {
	// String identifying the AMF Set ID (10 bits) as specified in clause 2.10.1 of 3GPP TS 23.003.  It is encoded as a string of 3 hexadecimal characters where the first character is limited to  values 0 to 3 (i.e. 10 bits).
	AmfSetId string `json:"amfSetId" yaml:"amfSetId" bson:"amfSetId,omitempty"`
	// String identifying the AMF Set ID (10 bits) as specified in clause 2.10.1 of 3GPP TS 23.003.  It is encoded as a string of 3 hexadecimal characters where the first character is limited to  values 0 to 3 (i.e. 10 bits)
	AmfRegionId             string              `json:"amfRegionId" yaml:"amfRegionId" bson:"amfRegionId,omitempty"`
	GuamiList               []Guami             `json:"guamiList" yaml:"guamiList" bson:"guamiList,omitempty"`
	TaiList                 []Tai               `json:"taiList,omitempty" yaml:"taiList" bson:"taiList,omitempty"`
	TaiRangeList            []TaiRange          `json:"taiRangeList,omitempty" yaml:"taiRangeList" bson:"taiRangeList,omitempty"`
	BackupInfoAmfFailure    []Guami             `json:"backupInfoAmfFailure,omitempty" yaml:"backupInfoAmfFailure" bson:"backupInfoAmfFailure,omitempty"`
	BackupInfoAmfRemoval    []Guami             `json:"backupInfoAmfRemoval,omitempty" yaml:"backupInfoAmfRemoval" bson:"backupInfoAmfRemoval,omitempty"`
	N2InterfaceAmfInfo      *N2InterfaceAmfInfo `json:"n2InterfaceAmfInfo,omitempty" yaml:"n2InterfaceAmfInfo" bson:"n2InterfaceAmfInfo,omitempty"`
	AmfOnboardingCapability bool                `json:"amfOnboardingCapability,omitempty" yaml:"amfOnboardingCapability" bson:"amfOnboardingCapability,omitempty"`
	HighLatencyCom          bool                `json:"highLatencyCom,omitempty" yaml:"highLatencyCom" bson:"highLatencyCom,omitempty"`
}
