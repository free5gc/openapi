/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type MultipleUnitUsage struct {
	RatingGroup          int32               `json:"ratingGroup" yaml:"ratingGroup" bson:"ratingGroup" mapstructure:"RatingGroup"`
	RequestedUnit        *RequestedUnit      `json:"requestedUnit,omitempty" yaml:"requestedUnit" bson:"requestedUnit" mapstructure:"RequestedUnit"`
	UsedUnitContainer    []UsedUnitContainer `json:"usedUnitContainer,omitempty" yaml:"usedUnitContainer" bson:"usedUnitContainer" mapstructure:"UsedUnitContainer"`
	UPFID                string              `json:"uPFID,omitempty" yaml:"uPFID" bson:"uPFID" mapstructure:"UPFID"`
	MultihomedPDUAddress *PduAddress         `json:"multihomedPDUAddress,omitempty" yaml:"multihomedPDUAddress" bson:"multihomedPDUAddress" mapstructure:"MultihomedPDUAddress"`
}
