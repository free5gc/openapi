/*
 * Nnef_PFDmanagement Service API
 *
 * Packet Flow Description Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * Source file: 3GPP TS 29.551 v17.9.0, 5G System; Packet Flow Description Management Service
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.551/
 *
 * API version: 1.2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	
)



// Contains information related to the access token request
type AccessTokenReq struct {
	Grant_type string `json:"grant_type,omitempty" yaml:"grant_type" bson:"grant_type,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId string `json:"nfInstanceId,omitempty" yaml:"nfInstanceId" bson:"nfInstanceId,omitempty"`
	NfType NFType `json:"nfType,omitempty" yaml:"nfType" bson:"nfType,omitempty"`
	TargetNfType NFType `json:"targetNfType,omitempty" yaml:"targetNfType" bson:"targetNfType,omitempty"`
	Scope string `json:"scope,omitempty" yaml:"scope" bson:"scope,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	TargetNfInstanceId string `json:"targetNfInstanceId,omitempty" yaml:"targetNfInstanceId" bson:"targetNfInstanceId,omitempty"`
	RequesterPlmn *PlmnId `json:"requesterPlmn,omitempty" yaml:"requesterPlmn" bson:"requesterPlmn,omitempty"`
	RequesterPlmnList []PlmnId `json:"requesterPlmnList,omitempty" yaml:"requesterPlmnList" bson:"requesterPlmnList,omitempty"`
	RequesterSnssaiList []Snssai `json:"requesterSnssaiList,omitempty" yaml:"requesterSnssaiList" bson:"requesterSnssaiList,omitempty"`
	// Fully Qualified Domain Name
	RequesterFqdn string `json:"requesterFqdn,omitempty" yaml:"requesterFqdn" bson:"requesterFqdn,omitempty"`
	RequesterSnpnList []PlmnIdNid `json:"requesterSnpnList,omitempty" yaml:"requesterSnpnList" bson:"requesterSnpnList,omitempty"`
	TargetPlmn *PlmnId `json:"targetPlmn,omitempty" yaml:"targetPlmn" bson:"targetPlmn,omitempty"`
	TargetSnpn *PlmnIdNid `json:"targetSnpn,omitempty" yaml:"targetSnpn" bson:"targetSnpn,omitempty"`
	TargetSnssaiList []Snssai `json:"targetSnssaiList,omitempty" yaml:"targetSnssaiList" bson:"targetSnssaiList,omitempty"`
	TargetNsiList []string `json:"targetNsiList,omitempty" yaml:"targetNsiList" bson:"targetNsiList,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	TargetNfSetId string `json:"targetNfSetId,omitempty" yaml:"targetNfSetId" bson:"targetNfSetId,omitempty"`
	// NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \"set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)   <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NID> encoded as defined in clause 5.4.2 (\"Nid\" data type definition)  <NFInstanceId> encoded as defined in clause 5.3.2  <ServiceName> encoded as defined in 3GPP TS 29.510  <Set ID> encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit. 
	TargetNfServiceSetId string `json:"targetNfServiceSetId,omitempty" yaml:"targetNfServiceSetId" bson:"targetNfServiceSetId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	HnrfAccessTokenUri string `json:"hnrfAccessTokenUri,omitempty" yaml:"hnrfAccessTokenUri" bson:"hnrfAccessTokenUri,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SourceNfInstanceId string `json:"sourceNfInstanceId,omitempty" yaml:"sourceNfInstanceId" bson:"sourceNfInstanceId,omitempty"`
}

