/*
 * Nchf_OfflineOnlyCharging
 *
 * OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291  V17.0.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents the Authorized Default QoS.
type AuthorizedDefaultQos struct {
	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255.
	Var5qi int32 `json:"5qi,omitempty" yaml:"5qi" bson:"5qi,omitempty"`
	Arp    *Arp  `json:"arp,omitempty" yaml:"arp" bson:"arp,omitempty"`
	// This data type is defined in the same way as the '5QiPriorityLevel' data type, but with the OpenAPI 'nullable: true' property.
	PriorityLevel int32 `json:"priorityLevel,omitempty" yaml:"priorityLevel" bson:"priorityLevel,omitempty"`
	// This data type is defined in the same way as the 'AverWindow' data type, but with the OpenAPI 'nullable: true' property.
	AverWindow int32 `json:"averWindow,omitempty" yaml:"averWindow" bson:"averWindow,omitempty"`
	// This data type is defined in the same way as the 'MaxDataBurstVol' data type, but with the OpenAPI 'nullable: true' property.
	MaxDataBurstVol int32 `json:"maxDataBurstVol,omitempty" yaml:"maxDataBurstVol" bson:"maxDataBurstVol,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	MaxbrUl string `json:"maxbrUl,omitempty" yaml:"maxbrUl" bson:"maxbrUl,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	MaxbrDl string `json:"maxbrDl,omitempty" yaml:"maxbrDl" bson:"maxbrDl,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	GbrUl string `json:"gbrUl,omitempty" yaml:"gbrUl" bson:"gbrUl,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	GbrDl string `json:"gbrDl,omitempty" yaml:"gbrDl" bson:"gbrDl,omitempty"`
	// This data type is defined in the same way as the 'ExtMaxDataBurstVol' data type, but with the OpenAPI 'nullable: true' property.
	ExtMaxDataBurstVol int32 `json:"extMaxDataBurstVol,omitempty" yaml:"extMaxDataBurstVol" bson:"extMaxDataBurstVol,omitempty"`
}
