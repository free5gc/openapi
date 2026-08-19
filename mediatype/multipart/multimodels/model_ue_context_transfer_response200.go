package multimodels

import (
	"errors"
	"reflect"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
)

type UEContextTransferResponse200 struct {
	JsonData                                *models.Amf_Comm_UeContextTransferRspData `json:"jsonData,omitempty"                    multipart:"contentType:application/json"          yaml:"jsonData"                    bson:"jsonData,omitempty"`
	UeRadioCapabilityN2Information          []byte                                    `json:"binaryDataN2Information,omitempty"     multipart:"contentType:application/vnd.3gpp.ngap" yaml:"binaryDataN2Information"     bson:"binaryDataN2Information,omitempty"`
	UeRadioCapabilityForPagingN2Information []byte                                    `json:"binaryDataN2InformationExt1,omitempty" multipart:"contentType:application/vnd.3gpp.ngap" yaml:"binaryDataN2InformationExt1" bson:"binaryDataN2InformationExt1,omitempty"`
	UeNbiotRadioCapabilityN2Information     []byte                                    `json:"binaryDataN2InformationExt2,omitempty" multipart:"contentType:application/vnd.3gpp.ngap" yaml:"binaryDataN2InformationExt2" bson:"binaryDataN2InformationExt2,omitempty"`
}

func (c *UEContextTransferResponse200) ReadFromModels(m *models.UEContextTransferResponse200) error {
	if m == nil {
		return errors.New("model is nil")
	}

	if m.JsonData == nil {
		return errors.New("jsonData is nil")
	}

	c.JsonData = m.JsonData

	n2DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.ngap")

	ueRadioCapability := c.JsonData.UeRadioCapability
	if ueRadioCapability != nil {
		contentId := ueRadioCapability.NgapData.ContentId
		if contentId == "" {
			return errors.New("ueRadioCapability.NgapData.ContentId is empty")
		}
		c.UeRadioCapabilityN2Information =
			multipart.MatchRelatedContent(contentId, n2DataFields)
	}

	ueRadioCapabilityForPaging := c.JsonData.UeRadioCapabilityForPaging
	if ueRadioCapabilityForPaging != nil {
		contentId := ueRadioCapabilityForPaging.NgapData.ContentId
		if contentId == "" {
			return errors.New("ueRadioCapabilityForPaging.NgapData.ContentId is empty")
		}
		c.UeRadioCapabilityForPagingN2Information =
			multipart.MatchRelatedContent(contentId, n2DataFields)
	}

	ueNbiotRadioCapability := c.JsonData.UeNbiotRadioCapability
	if ueNbiotRadioCapability != nil {
		contentId := ueNbiotRadioCapability.NgapData.ContentId
		if contentId == "" {
			return errors.New("ueNbiotRadioCapability.NgapData.ContentId is empty")
		}
		c.UeNbiotRadioCapabilityN2Information =
			multipart.MatchRelatedContent(contentId, n2DataFields)
	}

	return nil
}

func (c *UEContextTransferResponse200) BuildToModels() (*models.UEContextTransferResponse200, error) {
	if c == nil {
		return nil, errors.New("multipart model is nil")
	}

	if c.JsonData == nil {
		return nil, errors.New("jsonData is nil")
	}

	m := &models.UEContextTransferResponse200{}
	m.JsonData = c.JsonData

	n2DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.ngap")

	ueRadioCapability := c.JsonData.UeRadioCapability
	if ueRadioCapability != nil {
		contentId := ueRadioCapability.NgapData.ContentId
		if contentId == "" {
			return nil, errors.New("ueRadioCapability.NgapData.ContentId is empty")
		}
		n2DataFields[0].Set(reflect.ValueOf(&multipart.RelatedContent{
			ContentID: contentId,
			Content:   c.UeRadioCapabilityN2Information,
		}))
	}

	ueRadioCapabilityForPaging := c.JsonData.UeRadioCapabilityForPaging
	if ueRadioCapabilityForPaging != nil {
		contentId := ueRadioCapabilityForPaging.NgapData.ContentId
		if contentId == "" {
			return nil, errors.New("ueRadioCapabilityForPaging.NgapData.ContentId is empty")
		}
		n2DataFields[1].Set(reflect.ValueOf(&multipart.RelatedContent{
			ContentID: contentId,
			Content:   c.UeRadioCapabilityForPagingN2Information,
		}))
	}

	ueNbiotRadioCapability := c.JsonData.UeNbiotRadioCapability
	if ueNbiotRadioCapability != nil {
		contentId := ueNbiotRadioCapability.NgapData.ContentId
		if contentId == "" {
			return nil, errors.New("ueNbiotRadioCapability.NgapData.ContentId is empty")
		}
		n2DataFields[2].Set(reflect.ValueOf(&multipart.RelatedContent{
			ContentID: contentId,
			Content:   c.UeNbiotRadioCapabilityN2Information,
		}))
	}

	return m, nil
}
