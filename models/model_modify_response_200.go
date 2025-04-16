/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.13.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type ModifyResponse200 struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NfInstanceId        string `json:"nfInstanceId" yaml:"nfInstanceId" bson:"nfInstanceId,omitempty"`
	ImplicitUnsubscribe bool   `json:"implicitUnsubscribe,omitempty" yaml:"implicitUnsubscribe" bson:"implicitUnsubscribe,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expires *time.Time `json:"expires,omitempty" yaml:"expires" bson:"expires,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	CallbackReference     string      `json:"callbackReference" yaml:"callbackReference" bson:"callbackReference,omitempty"`
	AmfServiceName        ServiceName `json:"amfServiceName,omitempty" yaml:"amfServiceName" bson:"amfServiceName,omitempty"`
	MonitoredResourceUris []string    `json:"monitoredResourceUris" yaml:"monitoredResourceUris" bson:"monitoredResourceUris,omitempty"`
	SingleNssai           *Snssai     `json:"singleNssai,omitempty" yaml:"singleNssai" bson:"singleNssai,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn             string  `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	SubscriptionId  string  `json:"subscriptionId,omitempty" yaml:"subscriptionId" bson:"subscriptionId,omitempty"`
	PlmnId          *PlmnId `json:"plmnId,omitempty" yaml:"plmnId" bson:"plmnId,omitempty"`
	ImmediateReport bool    `json:"immediateReport,omitempty" yaml:"immediateReport" bson:"immediateReport,omitempty"`
	// The execution report contains an array of report items. Each report item indicates one  failed modification.
	Report []ReportItem `json:"report" yaml:"report" bson:"report,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures     string                       `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	ContextInfo           *ContextInfo                 `json:"contextInfo,omitempty" yaml:"contextInfo" bson:"contextInfo,omitempty"`
	NfChangeFilter        bool                         `json:"nfChangeFilter,omitempty" yaml:"nfChangeFilter" bson:"nfChangeFilter,omitempty"`
	UniqueSubscription    bool                         `json:"uniqueSubscription,omitempty" yaml:"uniqueSubscription" bson:"uniqueSubscription,omitempty"`
	ResetIds              []string                     `json:"resetIds,omitempty" yaml:"resetIds" bson:"resetIds,omitempty"`
	UeConSmfDataSubFilter *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty" yaml:"ueConSmfDataSubFilter" bson:"ueConSmfDataSubFilter,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	DataRestorationCallbackUri string `json:"dataRestorationCallbackUri,omitempty" yaml:"dataRestorationCallbackUri" bson:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              bool   `json:"udrRestartInd,omitempty" yaml:"udrRestartInd" bson:"udrRestartInd,omitempty"`
}
