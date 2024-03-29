/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type N3gaLocation struct {
	N3gppTai       *Tai              `json:"n3gppTai,omitempty" yaml:"n3gppTai" bson:"n3gppTai" mapstructure:"N3gppTai"`
	N3IwfId        string            `json:"n3IwfId,omitempty" yaml:"n3IwfId" bson:"n3IwfId" mapstructure:"N3IwfId"`
	UeIpv4Addr     string            `json:"ueIpv4Addr,omitempty" yaml:"ueIpv4Addr" bson:"ueIpv4Addr" mapstructure:"UeIpv4Addr"`
	UeIpv6Addr     string            `json:"ueIpv6Addr,omitempty" yaml:"ueIpv6Addr" bson:"ueIpv6Addr" mapstructure:"UeIpv6Addr"`
	PortNumber     int32             `json:"portNumber,omitempty" yaml:"portNumber" bson:"portNumber" mapstructure:"PortNumber"`
	TnapId         *TnapId           `json:"tnapId,omitempty" yaml:"tnapId" bson:"tnapId" mapstructure:"TnapId"`
	Protocol       TransportProtocol `json:"protocol,omitempty" yaml:"protocol" bson:"protocol" mapstructure:"Protocol"`
	TwapId         *TwapId           `json:"twapId,omitempty" yaml:"twapId" bson:"twapId" mapstructure:"TwapId"`
	HfcNodeId      *HfcNodeId        `json:"hfcNodeId,omitempty" yaml:"hfcNodeId" bson:"hfcNodeId" mapstructure:"HfcNodeId"`
	Gli            string            `json:"gli,omitempty" yaml:"gli" bson:"gli" mapstructure:"Gli"`
	W5gbanLineType LineType          `json:"w5gbanLineType,omitempty" yaml:"w5gbanLineType" bson:"w5gbanLineType" mapstructure:"W5gbanLineType"`
	Gci            string            `json:"gci,omitempty" yaml:"gci" bson:"gci" mapstructure:"Gci"`
}
