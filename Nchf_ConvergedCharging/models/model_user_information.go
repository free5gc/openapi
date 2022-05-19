/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UserInformation struct {
	ServedGPSI string `json:"servedGPSI,omitempty" yaml:"servedGPSI" bson:"servedGPSI" mapstructure:"ServedGPSI"`
	ServedPEI string `json:"servedPEI,omitempty" yaml:"servedPEI" bson:"servedPEI" mapstructure:"ServedPEI"`
	UnauthenticatedFlag bool `json:"unauthenticatedFlag,omitempty" yaml:"unauthenticatedFlag" bson:"unauthenticatedFlag" mapstructure:"UnauthenticatedFlag"`
	RoamerInOut RoamerInOut `json:"roamerInOut,omitempty" yaml:"roamerInOut" bson:"roamerInOut" mapstructure:"RoamerInOut"`
}
