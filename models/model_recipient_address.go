/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RecipientAddress struct {
	RecipientAddressInfo *SmAddressInfo `json:"recipientAddressInfo,omitempty" yaml:"recipientAddressInfo" bson:"recipientAddressInfo" mapstructure:"RecipientAddressInfo"`
	SMaddresseeType SmAddresseeType `json:"sMaddresseeType,omitempty" yaml:"sMaddresseeType" bson:"sMaddresseeType" mapstructure:"SMaddresseeType"`
}
