/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models
import (
	"time"
)

type PduSessionInformation struct {
	NetworkSlicingInfo *NetworkSlicingInfo `json:"networkSlicingInfo,omitempty" yaml:"networkSlicingInfo" bson:"networkSlicingInfo" mapstructure:"NetworkSlicingInfo"`
	PduSessionID int32 `json:"pduSessionID" yaml:"pduSessionID" bson:"pduSessionID" mapstructure:"PduSessionID"`
	PduType PduSessionType `json:"pduType,omitempty" yaml:"pduType" bson:"pduType" mapstructure:"PduType"`
	SscMode SscMode `json:"sscMode,omitempty" yaml:"sscMode" bson:"sscMode" mapstructure:"SscMode"`
	HPlmnId *PlmnId `json:"hPlmnId,omitempty" yaml:"hPlmnId" bson:"hPlmnId" mapstructure:"HPlmnId"`
	ServingNetworkFunctionID *ServingNetworkFunctionId `json:"servingNetworkFunctionID,omitempty" yaml:"servingNetworkFunctionID" bson:"servingNetworkFunctionID" mapstructure:"ServingNetworkFunctionID"`
	RatType RatType `json:"ratType,omitempty" yaml:"ratType" bson:"ratType" mapstructure:"RatType"`
	MAPDUNon3GPPRATType RatType `json:"mAPDUNon3GPPRATType,omitempty" yaml:"mAPDUNon3GPPRATType" bson:"mAPDUNon3GPPRATType" mapstructure:"MAPDUNon3GPPRATType"`
	DnnId string `json:"dnnId" yaml:"dnnId" bson:"dnnId" mapstructure:"DnnId"`
	DnnSelectionMode DnnSelectionMode `json:"dnnSelectionMode,omitempty" yaml:"dnnSelectionMode" bson:"dnnSelectionMode" mapstructure:"DnnSelectionMode"`
	ChargingCharacteristics string `json:"chargingCharacteristics,omitempty" yaml:"chargingCharacteristics" bson:"chargingCharacteristics" mapstructure:"ChargingCharacteristics"`
	ChargingCharacteristicsSelectionMode ChargingCharacteristicsSelectionMode `json:"chargingCharacteristicsSelectionMode,omitempty" yaml:"chargingCharacteristicsSelectionMode" bson:"chargingCharacteristicsSelectionMode" mapstructure:"ChargingCharacteristicsSelectionMode"`
	StartTime *time.Time `json:"startTime,omitempty" yaml:"startTime" bson:"startTime" mapstructure:"StartTime"`
	StopTime *time.Time `json:"stopTime,omitempty" yaml:"stopTime" bson:"stopTime" mapstructure:"StopTime"`
	Var3gppPSDataOffStatus Model3GpppsDataOffStatus `json:"3gppPSDataOffStatus,omitempty" yaml:"3gppPSDataOffStatus" bson:"3gppPSDataOffStatus" mapstructure:"Var3gppPSDataOffStatus"`
	SessionStopIndicator bool `json:"sessionStopIndicator,omitempty" yaml:"sessionStopIndicator" bson:"sessionStopIndicator" mapstructure:"SessionStopIndicator"`
	PduAddress *PduAddress `json:"pduAddress,omitempty" yaml:"pduAddress" bson:"pduAddress" mapstructure:"PduAddress"`
	Diagnostics int32 `json:"diagnostics,omitempty" yaml:"diagnostics" bson:"diagnostics" mapstructure:"Diagnostics"`
	AuthorizedQoSInformation *AuthorizedDefaultQos `json:"authorizedQoSInformation,omitempty" yaml:"authorizedQoSInformation" bson:"authorizedQoSInformation" mapstructure:"AuthorizedQoSInformation"`
	SubscribedQoSInformation *SubscribedDefaultQos `json:"subscribedQoSInformation,omitempty" yaml:"subscribedQoSInformation" bson:"subscribedQoSInformation" mapstructure:"SubscribedQoSInformation"`
	AuthorizedSessionAMBR *Ambr `json:"authorizedSessionAMBR,omitempty" yaml:"authorizedSessionAMBR" bson:"authorizedSessionAMBR" mapstructure:"AuthorizedSessionAMBR"`
	SubscribedSessionAMBR *Ambr `json:"subscribedSessionAMBR,omitempty" yaml:"subscribedSessionAMBR" bson:"subscribedSessionAMBR" mapstructure:"SubscribedSessionAMBR"`
	ServingCNPlmnId *PlmnId `json:"servingCNPlmnId,omitempty" yaml:"servingCNPlmnId" bson:"servingCNPlmnId" mapstructure:"ServingCNPlmnId"`
	MAPDUSessionInformation *MapduSessionInformation `json:"mAPDUSessionInformation,omitempty" yaml:"mAPDUSessionInformation" bson:"mAPDUSessionInformation" mapstructure:"MAPDUSessionInformation"`
	EnhancedDiagnostics []RanNasRelCause `json:"enhancedDiagnostics,omitempty" yaml:"enhancedDiagnostics" bson:"enhancedDiagnostics" mapstructure:"EnhancedDiagnostics"`
}
