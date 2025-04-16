/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V17.11.0; 5G System; Session Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Individual QoS flow to setup
type QosFlowSetupItem struct {
	// Unsigned integer identifying a QoS flow, within the range 0 to 63.
	Qfi int32 `json:"qfi" yaml:"qfi" bson:"qfi,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	QosRules string `json:"qosRules" yaml:"qosRules" bson:"qosRules,omitempty"`
	// EPS Bearer Identifier
	Ebi int32 `json:"ebi,omitempty" yaml:"ebi" bson:"ebi,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	QosFlowDescription string                       `json:"qosFlowDescription,omitempty" yaml:"qosFlowDescription" bson:"qosFlowDescription,omitempty"`
	QosFlowProfile     *SmfPduSessionQosFlowProfile `json:"qosFlowProfile,omitempty" yaml:"qosFlowProfile" bson:"qosFlowProfile,omitempty"`
	AssociatedAnType   QosFlowAccessType            `json:"associatedAnType,omitempty" yaml:"associatedAnType" bson:"associatedAnType,omitempty"`
	DefaultQosRuleInd  bool                         `json:"defaultQosRuleInd,omitempty" yaml:"defaultQosRuleInd" bson:"defaultQosRuleInd,omitempty"`
}
