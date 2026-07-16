package multimodels

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
)

type N1N2MessageTransferRequestBody struct {
	JsonData      *models.Amf_Comm_N1N2MessageTransferReqData `json:"jsonData,omitempty"                multipart:"contentType:application/json"           yaml:"jsonData"                bson:"jsonData,omitempty"`
	N1Message     []byte                                      `json:"binaryDataN1Message,omitempty"     multipart:"contentType:application/vnd.3gpp.5gnas" yaml:"binaryDataN1Message"     bson:"binaryDataN1Message,omitempty"`
	N2Information []byte                                      `json:"binaryDataN2Information,omitempty" multipart:"contentType:application/vnd.3gpp.ngap"  yaml:"binaryDataN2Information" bson:"binaryDataN2Information,omitempty"`
	MtData        []byte                                      `json:"binaryMtData,omitempty"            multipart:"contentType:application/vnd.3gpp.5gnas" yaml:"binaryMtData"            bson:"binaryMtData,omitempty"`
}

func (c *N1N2MessageTransferRequestBody) ReadFromModels(m *models.N1N2MessageTransferRequestBody) error {
	if m == nil {
		return errors.New("model is nil")
	}

	if m.JsonData == nil {
		return errors.New("jsonData is nil")
	}

	c.JsonData = m.JsonData

	n2DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.ngap")
	n1DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.5gnas")

	n1MsgContainer := c.JsonData.N1MessageContainer
	if n1MsgContainer != nil {
		if n1MsgContainer.N1MessageContent == nil {
			return errors.New("n1MessageContainer.N1MessageContent is nil")
		}
		contentId := n1MsgContainer.N1MessageContent.ContentId
		if contentId == "" {
			return errors.New("n1MessageContent.ContentId is empty")
		}
		c.N1Message = multipart.MatchRelatedContent(contentId, n1DataFields)
	}

	n2InfoContainer := c.JsonData.N2InfoContainer
	if n2InfoContainer != nil {
		var err error
		c.N2Information, err = readN2InformationFromRelatedContent(
			n2InfoContainer,
			n2DataFields,
		)
		if err != nil {
			return fmt.Errorf("failed to read N2Information: %w", err)
		}
	}

	if c.JsonData.MtData != nil {
		contentId := c.JsonData.MtData.ContentId
		if contentId == "" {
			return errors.New("mtData.ContentId is empty")
		}
		c.MtData = multipart.MatchRelatedContent(contentId, n1DataFields)
	}

	return nil
}

func (c *N1N2MessageTransferRequestBody) BuildToModels() (*models.N1N2MessageTransferRequestBody, error) {
	if c == nil {
		return nil, errors.New("multipart model is nil")
	}

	if c.JsonData == nil {
		return nil, errors.New("jsonData is nil")
	}

	m := &models.N1N2MessageTransferRequestBody{}
	m.JsonData = c.JsonData

	n1DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.5gnas")
	n2DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.ngap")
	n1FieldIndex := 0
	n2FieldIndex := 0

	n1MsgContainer := c.JsonData.N1MessageContainer
	if n1MsgContainer != nil && n1MsgContainer.N1MessageContent != nil {
		contentId := n1MsgContainer.N1MessageContent.ContentId
		if contentId == "" {
			return nil, errors.New("n1MessageContent.ContentId is empty")
		}
		n1DataFields[n1FieldIndex].Set(reflect.ValueOf(&multipart.RelatedContent{
			ContentID: contentId,
			Content:   c.N1Message,
		}))
		n1FieldIndex++
	}

	n2InfoContainer := c.JsonData.N2InfoContainer
	if n2InfoContainer != nil {
		err := putN2InformationToRelatedContent(n2InfoContainer, c.N2Information, n2DataFields[n2FieldIndex])
		if err != nil {
			return nil, fmt.Errorf("failed to put N2Information to related content: %w", err)
		}
	}
	if c.JsonData.MtData != nil {
		contentId := c.JsonData.MtData.ContentId
		if contentId == "" {
			return nil, errors.New("mtData.ContentId is empty")
		}
		n1DataFields[n1FieldIndex].Set(reflect.ValueOf(&multipart.RelatedContent{
			ContentID: contentId,
			Content:   c.MtData,
		}))
	}

	return m, nil
}
