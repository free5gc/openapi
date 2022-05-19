/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SmInterface struct {
	InterfaceId string `json:"interfaceId,omitempty" yaml:"interfaceId" bson:"interfaceId" mapstructure:"InterfaceId"`
	InterfaceText string `json:"interfaceText,omitempty" yaml:"interfaceText" bson:"interfaceText" mapstructure:"InterfaceText"`
	InterfacePort string `json:"interfacePort,omitempty" yaml:"interfacePort" bson:"interfacePort" mapstructure:"InterfacePort"`
	InterfaceType InterfaceType `json:"interfaceType,omitempty" yaml:"interfaceType" bson:"interfaceType" mapstructure:"InterfaceType"`
}