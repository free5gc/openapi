/*
 * Nhss_imsUECM
 *
 * Nhss UE Context Management Service for IMS. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * Source file: 3GPP TS 29.562 HSS Services for interworking with IMS, version 16.5.0
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.562/
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models



// Ims authorization request data
type AuthorizationRequest struct {
	Impi string `json:"impi,omitempty" yaml:"impi" bson:"impi"`
	AuthorizationType AuthorizationType `json:"authorizationType" yaml:"authorizationType" bson:"authorizationType"`
	VisitedNetworkIdentifier string `json:"visitedNetworkIdentifier,omitempty" yaml:"visitedNetworkIdentifier" bson:"visitedNetworkIdentifier"`
	EmergencyIndicator bool `json:"emergencyIndicator,omitempty" yaml:"emergencyIndicator" bson:"emergencyIndicator"`
	SupportedFeatures string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures"`
}