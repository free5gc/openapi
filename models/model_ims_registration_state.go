/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.562 HSS Services for interworking with IMS, version 17.6.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.562/
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ImsRegistrationState string

// List of ImsRegistrationState
const (
	ImsRegistrationState_REGISTERED                ImsRegistrationState = "REGISTERED"
	ImsRegistrationState_NOT_REGISTERED            ImsRegistrationState = "NOT_REGISTERED"
	ImsRegistrationState_AUTHENTICATION_PENDING    ImsRegistrationState = "AUTHENTICATION_PENDING"
	ImsRegistrationState_REGISTERED_UNREG_SERVICES ImsRegistrationState = "REGISTERED_UNREG_SERVICES"
)
