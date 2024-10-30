/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type CreateUeContextRequest struct {
	JsonData                     *UeContextCreateData `json:"jsonData,omitempty" yaml:"jsonData" bson:"jsonData,omitempty" multipart:"contentType:application/json,omitempty"`
	BinaryDataN2Information      []byte               `json:"binaryDataN2Information,omitempty" yaml:"binaryDataN2Information" bson:"binaryDataN2Information,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt1  []byte               `json:"binaryDataN2InformationExt1,omitempty" yaml:"binaryDataN2InformationExt1" bson:"binaryDataN2InformationExt1,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt2  []byte               `json:"binaryDataN2InformationExt2,omitempty" yaml:"binaryDataN2InformationExt2" bson:"binaryDataN2InformationExt2,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt3  []byte               `json:"binaryDataN2InformationExt3,omitempty" yaml:"binaryDataN2InformationExt3" bson:"binaryDataN2InformationExt3,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt4  []byte               `json:"binaryDataN2InformationExt4,omitempty" yaml:"binaryDataN2InformationExt4" bson:"binaryDataN2InformationExt4,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt5  []byte               `json:"binaryDataN2InformationExt5,omitempty" yaml:"binaryDataN2InformationExt5" bson:"binaryDataN2InformationExt5,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt6  []byte               `json:"binaryDataN2InformationExt6,omitempty" yaml:"binaryDataN2InformationExt6" bson:"binaryDataN2InformationExt6,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt7  []byte               `json:"binaryDataN2InformationExt7,omitempty" yaml:"binaryDataN2InformationExt7" bson:"binaryDataN2InformationExt7,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt8  []byte               `json:"binaryDataN2InformationExt8,omitempty" yaml:"binaryDataN2InformationExt8" bson:"binaryDataN2InformationExt8,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt9  []byte               `json:"binaryDataN2InformationExt9,omitempty" yaml:"binaryDataN2InformationExt9" bson:"binaryDataN2InformationExt9,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt10 []byte               `json:"binaryDataN2InformationExt10,omitempty" yaml:"binaryDataN2InformationExt10" bson:"binaryDataN2InformationExt10,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt11 []byte               `json:"binaryDataN2InformationExt11,omitempty" yaml:"binaryDataN2InformationExt11" bson:"binaryDataN2InformationExt11,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt12 []byte               `json:"binaryDataN2InformationExt12,omitempty" yaml:"binaryDataN2InformationExt12" bson:"binaryDataN2InformationExt12,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt13 []byte               `json:"binaryDataN2InformationExt13,omitempty" yaml:"binaryDataN2InformationExt13" bson:"binaryDataN2InformationExt13,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt14 []byte               `json:"binaryDataN2InformationExt14,omitempty" yaml:"binaryDataN2InformationExt14" bson:"binaryDataN2InformationExt14,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt15 []byte               `json:"binaryDataN2InformationExt15,omitempty" yaml:"binaryDataN2InformationExt15" bson:"binaryDataN2InformationExt15,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt16 []byte               `json:"binaryDataN2InformationExt16,omitempty" yaml:"binaryDataN2InformationExt16" bson:"binaryDataN2InformationExt16,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
	BinaryDataN2InformationExt17 []byte               `json:"binaryDataN2InformationExt17,omitempty" yaml:"binaryDataN2InformationExt17" bson:"binaryDataN2InformationExt17,omitempty" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.PduSessionList.N2InfoContent.NgapData.ContentId,omitempty"`
}
