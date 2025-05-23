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

type CreditManagementStatus string

// List of CreditManagementStatus
const (
	CreditManagementStatus_END_USER_SER_DENIED CreditManagementStatus = "END_USER_SER_DENIED" // #nosec G101
	CreditManagementStatus_CREDIT_CTRL_NOT_APP CreditManagementStatus = "CREDIT_CTRL_NOT_APP" // #nosec G101
	CreditManagementStatus_AUTH_REJECTED       CreditManagementStatus = "AUTH_REJECTED"
	CreditManagementStatus_USER_UNKNOWN        CreditManagementStatus = "USER_UNKNOWN"
	CreditManagementStatus_RATING_FAILED       CreditManagementStatus = "RATING_FAILED"
)
