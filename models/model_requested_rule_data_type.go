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

type RequestedRuleDataType string

// List of RequestedRuleDataType
const (
	RequestedRuleDataType_CH_ID         RequestedRuleDataType = "CH_ID"
	RequestedRuleDataType_MS_TIME_ZONE  RequestedRuleDataType = "MS_TIME_ZONE"
	RequestedRuleDataType_USER_LOC_INFO RequestedRuleDataType = "USER_LOC_INFO"
	RequestedRuleDataType_RES_RELEASE   RequestedRuleDataType = "RES_RELEASE"
	RequestedRuleDataType_SUCC_RES_ALLO RequestedRuleDataType = "SUCC_RES_ALLO"
	RequestedRuleDataType_EPS_FALLBACK  RequestedRuleDataType = "EPS_FALLBACK"
)
