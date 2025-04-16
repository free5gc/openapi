/*
 * AUSF API
 *
 * AUSF UE Authentication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.509 V17.9.0; 5G System; 3GPP TS Authentication Server services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.509
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type EapAuthMethodResponse200 struct {
	// contains an EAP packet
	EapPayload string `json:"eapPayload" yaml:"eapPayload" bson:"eapPayload,omitempty"`
	// URI : /{eapSessionUri}, a map(list of key-value pairs) where member serves as key
	Links map[string][]Link `json:"_links" yaml:"_links" bson:"_links,omitempty"`
}
