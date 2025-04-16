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

// Information of a CHF NF Instance
type ChfInfo struct {
	SupiRangeList []SupiRange     `json:"supiRangeList,omitempty" yaml:"supiRangeList" bson:"supiRangeList,omitempty"`
	GpsiRangeList []IdentityRange `json:"gpsiRangeList,omitempty" yaml:"gpsiRangeList" bson:"gpsiRangeList,omitempty"`
	PlmnRangeList []PlmnRange     `json:"plmnRangeList,omitempty" yaml:"plmnRangeList" bson:"plmnRangeList,omitempty"`
	// Identifier of a group of NFs.
	GroupId string `json:"groupId,omitempty" yaml:"groupId" bson:"groupId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	PrimaryChfInstance string `json:"primaryChfInstance,omitempty" yaml:"primaryChfInstance" bson:"primaryChfInstance,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	SecondaryChfInstance string `json:"secondaryChfInstance,omitempty" yaml:"secondaryChfInstance" bson:"secondaryChfInstance,omitempty"`
}
