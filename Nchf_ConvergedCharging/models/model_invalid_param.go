/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type InvalidParam struct {
	Param string `json:"param" yaml:"param" bson:"param" mapstructure:"Param"`
	Reason string `json:"reason,omitempty" yaml:"reason" bson:"reason" mapstructure:"Reason"`
}
