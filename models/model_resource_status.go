/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V17.11.0; 5G System; Session Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.2.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ResourceStatus string

// List of ResourceStatus
const (
	ResourceStatus_RELEASED       ResourceStatus = "RELEASED"
	ResourceStatus_UNCHANGED      ResourceStatus = "UNCHANGED"
	ResourceStatus_TRANSFERRED    ResourceStatus = "TRANSFERRED"
	ResourceStatus_UPDATED        ResourceStatus = "UPDATED"
	ResourceStatus_ALT_ANCHOR_SMF ResourceStatus = "ALT_ANCHOR_SMF"
)
