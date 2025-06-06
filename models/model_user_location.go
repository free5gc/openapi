/*
 * Nchf_OfflineOnlyCharging
 *
 * OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291  V17.0.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// At least one of eutraLocation, nrLocation and n3gaLocation shall be present. Several of them may be present.
type UserLocation struct {
	EutraLocation *EutraLocation `json:"eutraLocation,omitempty" yaml:"eutraLocation" bson:"eutraLocation,omitempty"`
	NrLocation    *NrLocation    `json:"nrLocation,omitempty" yaml:"nrLocation" bson:"nrLocation,omitempty"`
	N3gaLocation  *N3gaLocation  `json:"n3gaLocation,omitempty" yaml:"n3gaLocation" bson:"n3gaLocation,omitempty"`
	UtraLocation  *UtraLocation  `json:"utraLocation,omitempty" yaml:"utraLocation" bson:"utraLocation,omitempty"`
	GeraLocation  *GeraLocation  `json:"geraLocation,omitempty" yaml:"geraLocation" bson:"geraLocation,omitempty"`
}
