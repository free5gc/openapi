/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RedirectServer struct {
	RedirectAddressType   RedirectAddressType `json:"redirectAddressType" yaml:"redirectAddressType" bson:"redirectAddressType" mapstructure:"RedirectAddressType"`
	RedirectServerAddress string              `json:"redirectServerAddress" yaml:"redirectServerAddress" bson:"redirectServerAddress" mapstructure:"RedirectServerAddress"`
}
