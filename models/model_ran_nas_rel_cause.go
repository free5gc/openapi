/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RanNasRelCause struct {
	NgApCause    *NgApCause `json:"ngApCause,omitempty" yaml:"ngApCause" bson:"ngApCause" mapstructure:"NgApCause"`
	Var5gMmCause int32      `json:"5gMmCause,omitempty" yaml:"5gMmCause" bson:"5gMmCause" mapstructure:"Var5gMmCause"`
	Var5gSmCause int32      `json:"5gSmCause,omitempty" yaml:"5gSmCause" bson:"5gSmCause" mapstructure:"Var5gSmCause"`
	// Defines the EPS RAN/NAS release cause.
	EpsCause string `json:"epsCause,omitempty" yaml:"epsCause" bson:"epsCause" mapstructure:"EpsCause"`
}
