/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models
type DnnSelectionMode string

// List of dnnSelectionMode
const (
	DnnSelectionMode_VERIFIED DnnSelectionMode = "VERIFIED"
	DnnSelectionMode_UE_DNN_NOT_VERIFIED DnnSelectionMode = "UE_DNN_NOT_VERIFIED"
	DnnSelectionMode_NW_DNN_NOT_VERIFIED DnnSelectionMode = "NW_DNN_NOT_VERIFIED"
)