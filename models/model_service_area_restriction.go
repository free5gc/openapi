/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ServiceAreaRestriction struct {
	RestrictionType               RestrictionType `json:"restrictionType,omitempty" yaml:"restrictionType" bson:"restrictionType" mapstructure:"RestrictionType"`
	Areas                         []Area          `json:"areas,omitempty" yaml:"areas" bson:"areas" mapstructure:"Areas"`
	MaxNumOfTAs                   int32           `json:"maxNumOfTAs,omitempty" yaml:"maxNumOfTAs" bson:"maxNumOfTAs" mapstructure:"MaxNumOfTAs"`
	MaxNumOfTAsForNotAllowedAreas int32           `json:"maxNumOfTAsForNotAllowedAreas,omitempty" yaml:"maxNumOfTAsForNotAllowedAreas" bson:"maxNumOfTAsForNotAllowedAreas" mapstructure:"MaxNumOfTAsForNotAllowedAreas"`
}
