/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the Edge Configuration Server Address Configuration Information as defined in clause 5.2.3.6.1 of 3GPP TS 23.502.
type EcsServerAddr struct {
	EcsFqdnList      []string `json:"ecsFqdnList,omitempty" yaml:"ecsFqdnList" bson:"ecsFqdnList,omitempty"`
	EcsIpAddressList []IpAddr `json:"ecsIpAddressList,omitempty" yaml:"ecsIpAddressList" bson:"ecsIpAddressList,omitempty"`
	EcsUriList       []string `json:"ecsUriList,omitempty" yaml:"ecsUriList" bson:"ecsUriList,omitempty"`
	EcsProviderId    string   `json:"ecsProviderId,omitempty" yaml:"ecsProviderId" bson:"ecsProviderId,omitempty"`
}