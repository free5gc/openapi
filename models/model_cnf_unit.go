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

// During the processing of cnfUnits attribute, all the members in the array shall be  interpreted as logically concatenated with logical \"AND\".
type CnfUnit struct {
	CnfUnit []Atom `json:"cnfUnit" yaml:"cnfUnit" bson:"cnfUnit,omitempty"`
}
