/*
 * Nudm_UEAU
 *
 * UDM UE Authentication Service. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 16.6.0
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RgAuthCtx struct {
	AuthInd           bool   `json:"authInd" yaml:"authInd" bson:"authInd,omitempty"`
	Supi              string `json:"supi,omitempty" yaml:"supi" bson:"supi,omitempty"`
	SupportedFeatures string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
}
