/*
 * Nhss_imsUECM
 *
 * Nhss UE Context Management Service for IMS.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.562 Home Subscriber Server (HSS) Services, version 17.7.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.562/
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ImsRegistrationType string

// List of ImsRegistrationType
const (
	ImsRegistrationType_INITIAL_REGISTRATION          ImsRegistrationType = "INITIAL_REGISTRATION"
	ImsRegistrationType_RE_REGISTRATION               ImsRegistrationType = "RE_REGISTRATION"
	ImsRegistrationType_TIMEOUT_DEREGISTRATION        ImsRegistrationType = "TIMEOUT_DEREGISTRATION"
	ImsRegistrationType_USER_DEREGISTRATION           ImsRegistrationType = "USER_DEREGISTRATION"
	ImsRegistrationType_ADMINISTRATIVE_DEREGISTRATION ImsRegistrationType = "ADMINISTRATIVE_DEREGISTRATION"
	ImsRegistrationType_AUTHENTICATION_FAILURE        ImsRegistrationType = "AUTHENTICATION_FAILURE"
	ImsRegistrationType_AUTHENTICATION_TIMEOUT        ImsRegistrationType = "AUTHENTICATION_TIMEOUT"
	ImsRegistrationType_UNREGISTERED_USER             ImsRegistrationType = "UNREGISTERED_USER"
)
