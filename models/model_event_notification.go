/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.522 V17.7.0; 5G System; Network Exposure Function Northbound APIs.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.522/
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents a traffic influence event notification.
type EventNotification struct {
	// Identifies an NEF Northbound interface transaction, generated by the AF.
	AfTransId          string           `json:"afTransId,omitempty" yaml:"afTransId" bson:"afTransId,omitempty"`
	DnaiChgType        DnaiChangeType   `json:"dnaiChgType" yaml:"dnaiChgType" bson:"dnaiChgType,omitempty"`
	SourceTrafficRoute *RouteToLocation `json:"sourceTrafficRoute,omitempty" yaml:"sourceTrafficRoute" bson:"sourceTrafficRoute,omitempty"`
	SubscribedEvent    SubscribedEvent  `json:"subscribedEvent" yaml:"subscribedEvent" bson:"subscribedEvent,omitempty"`
	TargetTrafficRoute *RouteToLocation `json:"targetTrafficRoute,omitempty" yaml:"targetTrafficRoute" bson:"targetTrafficRoute,omitempty"`
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	SourceDnai string `json:"sourceDnai,omitempty" yaml:"sourceDnai" bson:"sourceDnai,omitempty"`
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	TargetDnai string `json:"targetDnai,omitempty" yaml:"targetDnai" bson:"targetDnai,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi string `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	SrcUeIpv4Addr   string `json:"srcUeIpv4Addr,omitempty" yaml:"srcUeIpv4Addr" bson:"srcUeIpv4Addr,omitempty"`
	SrcUeIpv6Prefix string `json:"srcUeIpv6Prefix,omitempty" yaml:"srcUeIpv6Prefix" bson:"srcUeIpv6Prefix,omitempty"`
	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	TgtUeIpv4Addr   string `json:"tgtUeIpv4Addr,omitempty" yaml:"tgtUeIpv4Addr" bson:"tgtUeIpv4Addr,omitempty"`
	TgtUeIpv6Prefix string `json:"tgtUeIpv6Prefix,omitempty" yaml:"tgtUeIpv6Prefix" bson:"tgtUeIpv6Prefix,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	UeMac string `json:"ueMac,omitempty" yaml:"ueMac" bson:"ueMac,omitempty"`
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	AfAckUri string `json:"afAckUri,omitempty" yaml:"afAckUri" bson:"afAckUri,omitempty"`
}
