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

type QfiContainerInformation struct {
	QFI int32 `json:"qFI,omitempty" yaml:"qFI" bson:"qFI" mapstructure:"QFI"`
	ReportTime *time.Time `json:"reportTime" yaml:"reportTime" bson:"reportTime" mapstructure:"ReportTime"`
	TimeofFirstUsage *time.Time `json:"timeofFirstUsage,omitempty" yaml:"timeofFirstUsage" bson:"timeofFirstUsage" mapstructure:"TimeofFirstUsage"`
	TimeofLastUsage *time.Time `json:"timeofLastUsage,omitempty" yaml:"timeofLastUsage" bson:"timeofLastUsage" mapstructure:"TimeofLastUsage"`
	QoSInformation *QosData `json:"qoSInformation,omitempty" yaml:"qoSInformation" bson:"qoSInformation" mapstructure:"QoSInformation"`
	QoSCharacteristics *QosCharacteristics `json:"qoSCharacteristics,omitempty" yaml:"qoSCharacteristics" bson:"qoSCharacteristics" mapstructure:"QoSCharacteristics"`
	UserLocationInformation *UserLocation `json:"userLocationInformation,omitempty" yaml:"userLocationInformation" bson:"userLocationInformation" mapstructure:"UserLocationInformation"`
	UetimeZone string `json:"uetimeZone,omitempty" yaml:"uetimeZone" bson:"uetimeZone" mapstructure:"UetimeZone"`
	PresenceReportingAreaInformation map[string]PresenceInfo `json:"presenceReportingAreaInformation,omitempty" yaml:"presenceReportingAreaInformation" bson:"presenceReportingAreaInformation" mapstructure:"PresenceReportingAreaInformation"`
	RATType RatType `json:"rATType,omitempty" yaml:"rATType" bson:"rATType" mapstructure:"RATType"`
	ServingNetworkFunctionID []ServingNetworkFunctionId `json:"servingNetworkFunctionID,omitempty" yaml:"servingNetworkFunctionID" bson:"servingNetworkFunctionID" mapstructure:"ServingNetworkFunctionID"`
	Var3gppPSDataOffStatus Model3GpppsDataOffStatus `json:"3gppPSDataOffStatus,omitempty" yaml:"3gppPSDataOffStatus" bson:"3gppPSDataOffStatus" mapstructure:"Var3gppPSDataOffStatus"`
	Var3gppChargingId int32 `json:"3gppChargingId,omitempty" yaml:"3gppChargingId" bson:"3gppChargingId" mapstructure:"Var3gppChargingId"`
	Diagnostics int32 `json:"diagnostics,omitempty" yaml:"diagnostics" bson:"diagnostics" mapstructure:"Diagnostics"`
	EnhancedDiagnostics []string `json:"enhancedDiagnostics,omitempty" yaml:"enhancedDiagnostics" bson:"enhancedDiagnostics" mapstructure:"EnhancedDiagnostics"`
}
