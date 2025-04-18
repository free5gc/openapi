/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.510 V17.12.0; 5G System; Network Function Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/
 *
 * API version: 1.2.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NoProfileMatchReason string

// List of NoProfileMatchReason
const (
	NoProfileMatchReason_REQUESTER_PLMN_NOT_ALLOWED        NoProfileMatchReason = "REQUESTER_PLMN_NOT_ALLOWED"
	NoProfileMatchReason_TARGET_NF_SUSPENDED               NoProfileMatchReason = "TARGET_NF_SUSPENDED"
	NoProfileMatchReason_TARGET_NF_UNDISCOVERABLE          NoProfileMatchReason = "TARGET_NF_UNDISCOVERABLE"
	NoProfileMatchReason_QUERY_PARAMS_COMBINATION_NO_MATCH NoProfileMatchReason = "QUERY_PARAMS_COMBINATION_NO_MATCH"
	NoProfileMatchReason_UNSPECIFIED                       NoProfileMatchReason = "UNSPECIFIED"
)
