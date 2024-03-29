/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains information related to the access token request
type AccessTokenReq struct {
	GrantType           string   `json:"grant_type" yaml:"grant_type" bson:"grant_type" mapstructure:"GrantType"`
	NfInstanceId        string   `json:"nfInstanceId" yaml:"nfInstanceId" bson:"nfInstanceId" mapstructure:"NfInstanceId"`
	NfType              NfType   `json:"nfType,omitempty" yaml:"nfType" bson:"nfType" mapstructure:"NfType"`
	TargetNfType        NfType   `json:"targetNfType,omitempty" yaml:"targetNfType" bson:"targetNfType" mapstructure:"TargetNfType"`
	Scope               string   `json:"scope" yaml:"scope" bson:"scope" mapstructure:"Scope"`
	TargetNfInstanceId  string   `json:"targetNfInstanceId,omitempty" yaml:"targetNfInstanceId" bson:"targetNfInstanceId" mapstructure:"TargetNfInstanceId"`
	RequesterPlmn       *PlmnId  `json:"requesterPlmn,omitempty" yaml:"requesterPlmn" bson:"requesterPlmn" mapstructure:"RequesterPlmn"`
	RequesterPlmnList   []PlmnId `json:"requesterPlmnList,omitempty" yaml:"requesterPlmnList" bson:"requesterPlmnList" mapstructure:"RequesterPlmnList"`
	RequesterSnssaiList []Snssai `json:"requesterSnssaiList,omitempty" yaml:"requesterSnssaiList" bson:"requesterSnssaiList" mapstructure:"RequesterSnssaiList"`
	// Fully Qualified Domain Name
	RequesterFqdn        string      `json:"requesterFqdn,omitempty" yaml:"requesterFqdn" bson:"requesterFqdn" mapstructure:"RequesterFqdn"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty" yaml:"requesterSnpnList" bson:"requesterSnpnList" mapstructure:"RequesterSnpnList"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty" yaml:"targetPlmn" bson:"targetPlmn" mapstructure:"TargetPlmn"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty" yaml:"targetSnssaiList" bson:"targetSnssaiList" mapstructure:"TargetSnssaiList"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty" yaml:"targetNsiList" bson:"targetNsiList" mapstructure:"TargetNsiList"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty" yaml:"targetNfSetId" bson:"targetNfSetId" mapstructure:"TargetNfSetId"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty" yaml:"targetNfServiceSetId" bson:"targetNfServiceSetId" mapstructure:"TargetNfServiceSetId"`
}
