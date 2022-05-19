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

type PduContainerInformation struct {
	TimeofFirstUsage *time.Time `json:"timeofFirstUsage,omitempty" yaml:"timeofFirstUsage" bson:"timeofFirstUsage" mapstructure:"TimeofFirstUsage"`
	TimeofLastUsage *time.Time `json:"timeofLastUsage,omitempty" yaml:"timeofLastUsage" bson:"timeofLastUsage" mapstructure:"TimeofLastUsage"`
	QoSInformation *QosData `json:"qoSInformation,omitempty" yaml:"qoSInformation" bson:"qoSInformation" mapstructure:"QoSInformation"`
	QoSCharacteristics *QosCharacteristics `json:"qoSCharacteristics,omitempty" yaml:"qoSCharacteristics" bson:"qoSCharacteristics" mapstructure:"QoSCharacteristics"`
	AfChargingIdentifier int32 `json:"afChargingIdentifier,omitempty" yaml:"afChargingIdentifier" bson:"afChargingIdentifier" mapstructure:"AfChargingIdentifier"`
	AfChargingIdString string `json:"afChargingIdString,omitempty" yaml:"afChargingIdString" bson:"afChargingIdString" mapstructure:"AfChargingIdString"`
	UserLocationInformation *UserLocation `json:"userLocationInformation,omitempty" yaml:"userLocationInformation" bson:"userLocationInformation" mapstructure:"UserLocationInformation"`
	UetimeZone string `json:"uetimeZone,omitempty" yaml:"uetimeZone" bson:"uetimeZone" mapstructure:"UetimeZone"`
	RATType RatType `json:"rATType,omitempty" yaml:"rATType" bson:"rATType" mapstructure:"RATType"`
	ServingNodeID []ServingNetworkFunctionId `json:"servingNodeID,omitempty" yaml:"servingNodeID" bson:"servingNodeID" mapstructure:"ServingNodeID"`
	PresenceReportingAreaInformation map[string]PresenceInfo `json:"presenceReportingAreaInformation,omitempty" yaml:"presenceReportingAreaInformation" bson:"presenceReportingAreaInformation" mapstructure:"PresenceReportingAreaInformation"`
	Var3gppPSDataOffStatus Model3GpppsDataOffStatus `json:"3gppPSDataOffStatus,omitempty" yaml:"3gppPSDataOffStatus" bson:"3gppPSDataOffStatus" mapstructure:"Var3gppPSDataOffStatus"`
	SponsorIdentity string `json:"sponsorIdentity,omitempty" yaml:"sponsorIdentity" bson:"sponsorIdentity" mapstructure:"SponsorIdentity"`
	ApplicationserviceProviderIdentity string `json:"applicationserviceProviderIdentity,omitempty" yaml:"applicationserviceProviderIdentity" bson:"applicationserviceProviderIdentity" mapstructure:"ApplicationserviceProviderIdentity"`
	ChargingRuleBaseName string `json:"chargingRuleBaseName,omitempty" yaml:"chargingRuleBaseName" bson:"chargingRuleBaseName" mapstructure:"ChargingRuleBaseName"`
	MAPDUSteeringFunctionality SteeringFunctionality `json:"mAPDUSteeringFunctionality,omitempty" yaml:"mAPDUSteeringFunctionality" bson:"mAPDUSteeringFunctionality" mapstructure:"MAPDUSteeringFunctionality"`
	MAPDUSteeringMode *SteeringMode `json:"mAPDUSteeringMode,omitempty" yaml:"mAPDUSteeringMode" bson:"mAPDUSteeringMode" mapstructure:"MAPDUSteeringMode"`
}