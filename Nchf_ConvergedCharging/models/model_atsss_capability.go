/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AtsssCapability struct {
	AtsssLL bool `json:"atsssLL,omitempty" yaml:"atsssLL" bson:"atsssLL" mapstructure:"AtsssLL"`
	Mptcp bool `json:"mptcp,omitempty" yaml:"mptcp" bson:"mptcp" mapstructure:"Mptcp"`
	RttWithoutPmf bool `json:"rttWithoutPmf,omitempty" yaml:"rttWithoutPmf" bson:"rttWithoutPmf" mapstructure:"RttWithoutPmf"`
}
