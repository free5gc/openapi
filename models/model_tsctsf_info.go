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

// Information of a TSCTSF NF Instance
type TsctsfInfo struct {
	// A map (list of key-value pairs) where a valid JSON string serves as key
	SNssaiInfoList                 map[string]SnssaiTsctsfInfoItem `json:"sNssaiInfoList,omitempty" yaml:"sNssaiInfoList" bson:"sNssaiInfoList,omitempty"`
	ExternalGroupIdentifiersRanges []IdentityRange                 `json:"externalGroupIdentifiersRanges,omitempty" yaml:"externalGroupIdentifiersRanges" bson:"externalGroupIdentifiersRanges,omitempty"`
	SupiRanges                     []SupiRange                     `json:"supiRanges,omitempty" yaml:"supiRanges" bson:"supiRanges,omitempty"`
	GpsiRanges                     []IdentityRange                 `json:"gpsiRanges,omitempty" yaml:"gpsiRanges" bson:"gpsiRanges,omitempty"`
	InternalGroupIdentifiersRanges []InternalGroupIdRange          `json:"internalGroupIdentifiersRanges,omitempty" yaml:"internalGroupIdentifiersRanges" bson:"internalGroupIdentifiersRanges,omitempty"`
}