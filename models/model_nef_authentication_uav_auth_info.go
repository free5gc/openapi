/*
 * Nnef_Authentication
 *
 * NEF Auth Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.256 V17.3.0; 5G System;Uncrewed Aerial Systems Network Function (UAS-NF); Aerial Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.256/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// UAV auth data
type NefAuthenticationUavAuthInfo struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi           string `json:"gpsi" yaml:"gpsi" bson:"gpsi,omitempty"`
	ServiceLevelId string `json:"serviceLevelId" yaml:"serviceLevelId" bson:"serviceLevelId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	AuthNotificationURI string  `json:"authNotificationURI,omitempty" yaml:"authNotificationURI" bson:"authNotificationURI,omitempty"`
	IpAddr              *IpAddr `json:"ipAddr,omitempty" yaml:"ipAddr" bson:"ipAddr,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.
	Pei               string                           `json:"pei,omitempty" yaml:"pei" bson:"pei,omitempty"`
	AuthServerAddress string                           `json:"authServerAddress,omitempty" yaml:"authServerAddress" bson:"authServerAddress,omitempty"`
	AuthMsg           *RefToBinaryData                 `json:"authMsg,omitempty" yaml:"authMsg" bson:"authMsg,omitempty"`
	AuthContainer     []NefAuthenticationAuthContainer `json:"authContainer,omitempty" yaml:"authContainer" bson:"authContainer,omitempty"`
	UeLocInfo         *UserLocation                    `json:"ueLocInfo,omitempty" yaml:"ueLocInfo" bson:"ueLocInfo,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn    string                `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	SNssai *ExtSnssai            `json:"sNssai,omitempty" yaml:"sNssai" bson:"sNssai,omitempty"`
	NfType NrfNfManagementNfType `json:"nfType" yaml:"nfType" bson:"nfType,omitempty"`
}
