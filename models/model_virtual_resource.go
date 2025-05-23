/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291 V17.9.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 3.1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type VirtualResource struct {
	VirtualMemory int32  `json:"virtualMemory,omitempty" yaml:"virtualMemory" bson:"virtualMemory,omitempty"`
	VirtualDisk   int32  `json:"virtualDisk,omitempty" yaml:"virtualDisk" bson:"virtualDisk,omitempty"`
	VirutalCPU    string `json:"virutalCPU,omitempty" yaml:"virutalCPU" bson:"virutalCPU,omitempty"`
}
