/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ImmediateReport struct {
	AmData                          *AccessAndMobilitySubscriptionData `json:"amData,omitempty" yaml:"amData" bson:"amData,omitempty"`
	SmfSelData                      *SmfSelectionSubscriptionData      `json:"smfSelData,omitempty" yaml:"smfSelData" bson:"smfSelData,omitempty"`
	UecAmfData                      *UeContextInAmfData                `json:"uecAmfData,omitempty" yaml:"uecAmfData" bson:"uecAmfData,omitempty"`
	UecSmfData                      *UeContextInSmfData                `json:"uecSmfData,omitempty" yaml:"uecSmfData" bson:"uecSmfData,omitempty"`
	UecSmsfData                     *UeContextInSmsfData               `json:"uecSmsfData,omitempty" yaml:"uecSmsfData" bson:"uecSmsfData,omitempty"`
	SmsSubsData                     *SmsSubscriptionData               `json:"smsSubsData,omitempty" yaml:"smsSubsData" bson:"smsSubsData,omitempty"`
	SmData                          *SmSubsData                        `json:"smData,omitempty" yaml:"smData" bson:"smData,omitempty"`
	TraceData                       *TraceData                         `json:"traceData,omitempty" yaml:"traceData" bson:"traceData,omitempty"`
	SmsMngData                      *SmsManagementSubscriptionData     `json:"smsMngData,omitempty" yaml:"smsMngData" bson:"smsMngData,omitempty"`
	LcsPrivacyData                  *LcsPrivacyData                    `json:"lcsPrivacyData,omitempty" yaml:"lcsPrivacyData" bson:"lcsPrivacyData,omitempty"`
	LcsMoData                       *LcsMoData                         `json:"lcsMoData,omitempty" yaml:"lcsMoData" bson:"lcsMoData,omitempty"`
	V2xData                         *V2xSubscriptionData               `json:"v2xData,omitempty" yaml:"v2xData" bson:"v2xData,omitempty"`
	LcsBroadcastAssistanceTypesData *LcsBroadcastAssistanceTypesData   `json:"lcsBroadcastAssistanceTypesData,omitempty" yaml:"lcsBroadcastAssistanceTypesData" bson:"lcsBroadcastAssistanceTypesData,omitempty"`
	ProseData                       *ProseSubscriptionData             `json:"proseData,omitempty" yaml:"proseData" bson:"proseData,omitempty"`
	MbsData                         *MbsSubscriptionData               `json:"mbsData,omitempty" yaml:"mbsData" bson:"mbsData,omitempty"`
	UcData                          *UcSubscriptionData                `json:"ucData,omitempty" yaml:"ucData" bson:"ucData,omitempty"`
}
