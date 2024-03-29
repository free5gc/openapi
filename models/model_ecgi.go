/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Ecgi struct {
	PlmnId      *PlmnId `json:"plmnId" yaml:"plmnId" bson:"plmnId" mapstructure:"PlmnId"`
	EutraCellId string  `json:"eutraCellId" yaml:"eutraCellId" bson:"eutraCellId" mapstructure:"EutraCellId"`
	Nid         string  `json:"nid,omitempty" yaml:"nid" bson:"nid" mapstructure:"Nid"`
}
