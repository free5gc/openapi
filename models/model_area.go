/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Area struct {
	Tacs     []string `json:"tacs,omitempty" yaml:"tacs" bson:"tacs" mapstructure:"Tacs"`
	AreaCode string   `json:"areaCode,omitempty" yaml:"areaCode" bson:"areaCode" mapstructure:"AreaCode"`
}
