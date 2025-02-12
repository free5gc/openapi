/*
 * Unified Data Repository Service API file for policy data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// DnnRouteSelectionDescriptor - Contains the route selector parameters (PDU session types, SSC modes and ATSSS information) per DNN
type DnnRouteSelectionDescriptor struct {

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn"`

	SscModes []SscMode `json:"sscModes,omitempty"`

	PduSessTypes []PduSessionType `json:"pduSessTypes,omitempty"`

	// Indicates whether MA PDU session establishment is allowed for this DNN. When set to value true MA PDU session establishment is allowed for this DNN.
	AtsssInfo bool `json:"atsssInfo,omitempty"`
}
