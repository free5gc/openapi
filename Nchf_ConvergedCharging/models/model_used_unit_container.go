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

type UsedUnitContainer struct {
	ServiceId int32 `json:"serviceId,omitempty" yaml:"serviceId" bson:"serviceId" mapstructure:"ServiceId"`
	QuotaManagementIndicator QuotaManagementIndicator `json:"quotaManagementIndicator,omitempty" yaml:"quotaManagementIndicator" bson:"quotaManagementIndicator" mapstructure:"QuotaManagementIndicator"`
	Triggers []Trigger `json:"triggers,omitempty" yaml:"triggers" bson:"triggers" mapstructure:"Triggers"`
	TriggerTimestamp *time.Time `json:"triggerTimestamp,omitempty" yaml:"triggerTimestamp" bson:"triggerTimestamp" mapstructure:"TriggerTimestamp"`
	Time int32 `json:"time,omitempty" yaml:"time" bson:"time" mapstructure:"Time"`
	TotalVolume int32 `json:"totalVolume,omitempty" yaml:"totalVolume" bson:"totalVolume" mapstructure:"TotalVolume"`
	UplinkVolume int32 `json:"uplinkVolume,omitempty" yaml:"uplinkVolume" bson:"uplinkVolume" mapstructure:"UplinkVolume"`
	DownlinkVolume int32 `json:"downlinkVolume,omitempty" yaml:"downlinkVolume" bson:"downlinkVolume" mapstructure:"DownlinkVolume"`
	ServiceSpecificUnits int32 `json:"serviceSpecificUnits,omitempty" yaml:"serviceSpecificUnits" bson:"serviceSpecificUnits" mapstructure:"ServiceSpecificUnits"`
	EventTimeStamps []*time.Time`json:"eventTimeStamps,omitempty" yaml:"eventTimeStamps" bson:"eventTimeStamps" mapstructure:"EventTimeStamps"`
	LocalSequenceNumber int32 `json:"localSequenceNumber" yaml:"localSequenceNumber" bson:"localSequenceNumber" mapstructure:"LocalSequenceNumber"`
	PDUContainerInformation *PduContainerInformation `json:"pDUContainerInformation,omitempty" yaml:"pDUContainerInformation" bson:"pDUContainerInformation" mapstructure:"PDUContainerInformation"`
	NSPAContainerInformation *NspaContainerInformation `json:"nSPAContainerInformation,omitempty" yaml:"nSPAContainerInformation" bson:"nSPAContainerInformation" mapstructure:"NSPAContainerInformation"`
}