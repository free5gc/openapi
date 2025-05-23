/*
 * Npcf_EventExposure
 *
 * PCF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.523 V17.7.0; 5G System; Policy Control Event Exposure Service; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.523/
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PcEvent string

// List of PcEvent
const (
	PcEvent_AC_TY_CH                PcEvent = "AC_TY_CH"
	PcEvent_PLMN_CH                 PcEvent = "PLMN_CH"
	PcEvent_SAC_CH                  PcEvent = "SAC_CH"
	PcEvent_SAT_CATEGORY_CH         PcEvent = "SAT_CATEGORY_CH"
	PcEvent_SUCCESS_UE_POL_DEL_SP   PcEvent = "SUCCESS_UE_POL_DEL_SP"
	PcEvent_UNSUCCESS_UE_POL_DEL_SP PcEvent = "UNSUCCESS_UE_POL_DEL_SP"
)
