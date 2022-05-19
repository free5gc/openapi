/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type InvocationResult struct {
	Error *ProblemDetails `json:"error,omitempty" yaml:"error" bson:"error" mapstructure:"Error"`
	FailureHandling FailureHandling `json:"failureHandling,omitempty" yaml:"failureHandling" bson:"failureHandling" mapstructure:"FailureHandling"`
}
