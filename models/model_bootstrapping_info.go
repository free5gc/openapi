/*
 * NRF Bootstrapping
 *
 * NRF Bootstrapping.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.510 V17.6.0; 5G System; Network Function Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Information returned by NRF in the bootstrapping response message
type BootstrappingInfo struct {
	Status Status `json:"status,omitempty" yaml:"status" bson:"status,omitempty"`
	// Map of link objects where the keys are the link relations defined in 3GPP TS 29.510 clause 6.4.6.3.3
	Links map[string][]Link `json:"_links" yaml:"_links" bson:"_links,omitempty"`
	// Map of features supported by the NRF, where the keys are the NRF services as defined in 3GPP TS 29.510 clause 6.1.6.3.11
	NrfFeatures map[string]string `json:"nrfFeatures,omitempty" yaml:"nrfFeatures" bson:"nrfFeatures,omitempty"`
	// Map indicating whether the NRF requires Oauth2-based authorization for accessing its services. The key of the map shall be the name of an NRF service, e.g. \"nnrf-nfm\" or \"nnrf-disc\"
	Oauth2Required map[string]bool `json:"oauth2Required,omitempty" yaml:"oauth2Required" bson:"oauth2Required,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.
	NrfSetId string `json:"nrfSetId,omitempty" yaml:"nrfSetId" bson:"nrfSetId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NrfInstanceId string `json:"nrfInstanceId,omitempty" yaml:"nrfInstanceId" bson:"nrfInstanceId,omitempty"`
}
