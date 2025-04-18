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

// Information of a NWDAF NF Instance
type NwdafInfo struct {
	EventIds        []EventId        `json:"eventIds,omitempty" yaml:"eventIds" bson:"eventIds,omitempty"`
	NwdafEvents     []NwdafEvent     `json:"nwdafEvents,omitempty" yaml:"nwdafEvents" bson:"nwdafEvents,omitempty"`
	TaiList         []Tai            `json:"taiList,omitempty" yaml:"taiList" bson:"taiList,omitempty"`
	TaiRangeList    []TaiRange       `json:"taiRangeList,omitempty" yaml:"taiRangeList" bson:"taiRangeList,omitempty"`
	NwdafCapability *NwdafCapability `json:"nwdafCapability,omitempty" yaml:"nwdafCapability" bson:"nwdafCapability,omitempty"`
	// indicating a time in seconds.
	AnalyticsDelay     int32                   `json:"analyticsDelay,omitempty" yaml:"analyticsDelay" bson:"analyticsDelay,omitempty"`
	ServingNfSetIdList []string                `json:"servingNfSetIdList,omitempty" yaml:"servingNfSetIdList" bson:"servingNfSetIdList,omitempty"`
	ServingNfTypeList  []NrfNfManagementNfType `json:"servingNfTypeList,omitempty" yaml:"servingNfTypeList" bson:"servingNfTypeList,omitempty"`
	MlAnalyticsList    []MlAnalyticsInfo       `json:"mlAnalyticsList,omitempty" yaml:"mlAnalyticsList" bson:"mlAnalyticsList,omitempty"`
}
