/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.512 V16.10.0; 5G System; Session Management Policy Control Service.
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.512/
 *
 * API version: 1.1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NullResponse200 struct {
	AccessType       AccessType          `json:"accessType,omitempty" yaml:"accessType" bson:"accessType,omitempty"`
	RatType          RatType             `json:"ratType,omitempty" yaml:"ratType" bson:"ratType,omitempty"`
	ServNfId         *ServingNfIdentity  `json:"servNfId,omitempty" yaml:"servNfId" bson:"servNfId,omitempty"`
	ServingNetwork   *PlmnIdNid          `json:"servingNetwork,omitempty" yaml:"servingNetwork" bson:"servingNetwork,omitempty"`
	UserLocationInfo *UserLocation       `json:"userLocationInfo,omitempty" yaml:"userLocationInfo" bson:"userLocationInfo,omitempty"`
	UeTimeZone       string              `json:"ueTimeZone,omitempty" yaml:"ueTimeZone" bson:"ueTimeZone,omitempty"`
	NetLocAccSupp    NetLocAccessSupport `json:"netLocAccSupp,omitempty" yaml:"netLocAccSupp" bson:"netLocAccSupp,omitempty"`
}