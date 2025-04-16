/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.514 V17.9.0; 5G System; Policy Authorization Service; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.514/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Identifies a media component.
type MediaComponent struct {
	// Contains an AF application identifier.
	AfAppId      string                `json:"afAppId,omitempty" yaml:"afAppId" bson:"afAppId,omitempty"`
	AfRoutReq    *AfRoutingRequirement `json:"afRoutReq,omitempty" yaml:"afRoutReq" bson:"afRoutReq,omitempty"`
	QosReference string                `json:"qosReference,omitempty" yaml:"qosReference" bson:"qosReference,omitempty"`
	DisUeNotif   bool                  `json:"disUeNotif,omitempty" yaml:"disUeNotif" bson:"disUeNotif,omitempty"`
	AltSerReqs   []string              `json:"altSerReqs,omitempty" yaml:"altSerReqs" bson:"altSerReqs,omitempty"`
	// Contains alternative service requirements that include individual QoS parameter sets.
	AltSerReqsData []AlternativeServiceRequirementsData `json:"altSerReqsData,omitempty" yaml:"altSerReqsData" bson:"altSerReqsData,omitempty"`
	// Represents the content version of some content.
	ContVer int32    `json:"contVer,omitempty" yaml:"contVer" bson:"contVer,omitempty"`
	Codecs  []string `json:"codecs,omitempty" yaml:"codecs" bson:"codecs,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	DesMaxLatency float32 `json:"desMaxLatency,omitempty" yaml:"desMaxLatency" bson:"desMaxLatency,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	DesMaxLoss float32    `json:"desMaxLoss,omitempty" yaml:"desMaxLoss" bson:"desMaxLoss,omitempty"`
	FlusId     string     `json:"flusId,omitempty" yaml:"flusId" bson:"flusId,omitempty"`
	FStatus    FlowStatus `json:"fStatus,omitempty" yaml:"fStatus" bson:"fStatus,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwDl string `json:"marBwDl,omitempty" yaml:"marBwDl" bson:"marBwDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwUl string `json:"marBwUl,omitempty" yaml:"marBwUl" bson:"marBwUl,omitempty"`
	// This data type is defined in the same way as the 'PacketLossRate' data type, but with the OpenAPI 'nullable: true' property.
	MaxPacketLossRateDl int32 `json:"maxPacketLossRateDl,omitempty" yaml:"maxPacketLossRateDl" bson:"maxPacketLossRateDl,omitempty"`
	// This data type is defined in the same way as the 'PacketLossRate' data type, but with the OpenAPI 'nullable: true' property.
	MaxPacketLossRateUl int32 `json:"maxPacketLossRateUl,omitempty" yaml:"maxPacketLossRateUl" bson:"maxPacketLossRateUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MaxSuppBwDl string `json:"maxSuppBwDl,omitempty" yaml:"maxSuppBwDl" bson:"maxSuppBwDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MaxSuppBwUl string `json:"maxSuppBwUl,omitempty" yaml:"maxSuppBwUl" bson:"maxSuppBwUl,omitempty"`
	MedCompN    int32  `json:"medCompN" yaml:"medCompN" bson:"medCompN,omitempty"`
	// Contains the requested bitrate and filters for the set of service data flows identified by their common flow identifier. The key of the map is the fNum attribute.
	MedSubComps map[string]MediaSubComponent `json:"medSubComps,omitempty" yaml:"medSubComps" bson:"medSubComps,omitempty"`
	MedType     MediaType                    `json:"medType,omitempty" yaml:"medType" bson:"medType,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MinDesBwDl string `json:"minDesBwDl,omitempty" yaml:"minDesBwDl" bson:"minDesBwDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MinDesBwUl string `json:"minDesBwUl,omitempty" yaml:"minDesBwUl" bson:"minDesBwUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MirBwDl string `json:"mirBwDl,omitempty" yaml:"mirBwDl" bson:"mirBwDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MirBwUl        string                   `json:"mirBwUl,omitempty" yaml:"mirBwUl" bson:"mirBwUl,omitempty"`
	PreemptCap     PreemptionCapability     `json:"preemptCap,omitempty" yaml:"preemptCap" bson:"preemptCap,omitempty"`
	PreemptVuln    PreemptionVulnerability  `json:"preemptVuln,omitempty" yaml:"preemptVuln" bson:"preemptVuln,omitempty"`
	PrioSharingInd PrioritySharingIndicator `json:"prioSharingInd,omitempty" yaml:"prioSharingInd" bson:"prioSharingInd,omitempty"`
	ResPrio        ReservPriority           `json:"resPrio,omitempty" yaml:"resPrio" bson:"resPrio,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	RrBw string `json:"rrBw,omitempty" yaml:"rrBw" bson:"rrBw,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	RsBw string `json:"rsBw,omitempty" yaml:"rsBw" bson:"rsBw,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	SharingKeyDl int32 `json:"sharingKeyDl,omitempty" yaml:"sharingKeyDl" bson:"sharingKeyDl,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	SharingKeyUl int32                `json:"sharingKeyUl,omitempty" yaml:"sharingKeyUl" bson:"sharingKeyUl,omitempty"`
	TsnQos       *TsnQosContainer     `json:"tsnQos,omitempty" yaml:"tsnQos" bson:"tsnQos,omitempty"`
	TscaiInputDl *TscaiInputContainer `json:"tscaiInputDl,omitempty" yaml:"tscaiInputDl" bson:"tscaiInputDl,omitempty"`
	TscaiInputUl *TscaiInputContainer `json:"tscaiInputUl,omitempty" yaml:"tscaiInputUl" bson:"tscaiInputUl,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TscaiTimeDom int32 `json:"tscaiTimeDom,omitempty" yaml:"tscaiTimeDom" bson:"tscaiTimeDom,omitempty"`
}
