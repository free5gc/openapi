/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.512 V17.11.0; 5G System; Session Management Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// describes the address of the SGSN
type SgsnAddress struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	SgsnIpv4Addr string `json:"sgsnIpv4Addr,omitempty" yaml:"sgsnIpv4Addr" bson:"sgsnIpv4Addr,omitempty"`
	SgsnIpv6Addr string `json:"sgsnIpv6Addr,omitempty" yaml:"sgsnIpv6Addr" bson:"sgsnIpv6Addr,omitempty"`
}
