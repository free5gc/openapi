/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type LocationReportingChargingInformation struct {
	LocationReportingMessageType     int32                   `json:"locationReportingMessageType" yaml:"locationReportingMessageType" bson:"locationReportingMessageType" mapstructure:"LocationReportingMessageType"`
	UserInformation                  *UserInformation        `json:"userInformation,omitempty" yaml:"userInformation" bson:"userInformation" mapstructure:"UserInformation"`
	UserLocationinfo                 *UserLocation           `json:"userLocationinfo,omitempty" yaml:"userLocationinfo" bson:"userLocationinfo" mapstructure:"UserLocationinfo"`
	PSCellInformation                *PsCellInformation      `json:"pSCellInformation,omitempty" yaml:"pSCellInformation" bson:"pSCellInformation" mapstructure:"PSCellInformation"`
	UetimeZone                       string                  `json:"uetimeZone,omitempty" yaml:"uetimeZone" bson:"uetimeZone" mapstructure:"UetimeZone"`
	RATType                          RatType                 `json:"rATType,omitempty" yaml:"rATType" bson:"rATType" mapstructure:"RATType"`
	PresenceReportingAreaInformation map[string]PresenceInfo `json:"presenceReportingAreaInformation,omitempty" yaml:"presenceReportingAreaInformation" bson:"presenceReportingAreaInformation" mapstructure:"PresenceReportingAreaInformation"`
}
