package multimodels

import (
	"errors"
	"reflect"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
)

type N1MessageNotifyRequestBody struct {
	JsonData            *models.Amf_Comm_N1MessageNotification `json:"jsonData,omitempty"            multipart:"contentType:application/json"           yaml:"jsonData"            bson:"jsonData,omitempty"`
	BinaryDataN1Message []byte                                 `json:"binaryDataN1Message,omitempty" multipart:"contentType:application/vnd.3gpp.5gnas" yaml:"binaryDataN1Message" bson:"binaryDataN1Message,omitempty"`
}

func (c *N1MessageNotifyRequestBody) ReadFromModels(m *models.N1MessageNotifyRequestBody) error {
	if m == nil {
		return errors.New("model is nil")
	}

	c.JsonData = m.JsonData
	if c.JsonData == nil {
		return errors.New("jsonData is nil")
	}

	n1MsgContainer := c.JsonData.N1MessageContainer
	if n1MsgContainer != nil {
		if n1MsgContainer.N1MessageContent == nil {
			return errors.New("n1MessageContainer.N1MessageContent is nil")
		}
		contentId := n1MsgContainer.N1MessageContent.ContentId
		if contentId == "" {
			return errors.New("n1MessageContainer.N1MessageContent.ContentId is empty")
		}
		n1DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.5gnas")
		c.BinaryDataN1Message = multipart.MatchRelatedContent(contentId, n1DataFields)
	}
	return nil
}

func (c *N1MessageNotifyRequestBody) BuildToModels() (*models.N1MessageNotifyRequestBody, error) {
	if c == nil {
		return nil, errors.New("multipart model is nil")
	}

	if c.JsonData == nil {
		return nil, errors.New("jsonData is nil")
	}

	m := &models.N1MessageNotifyRequestBody{}
	m.JsonData = c.JsonData

	if m.JsonData.N1MessageContainer != nil {
		n1MsgContainer := c.JsonData.N1MessageContainer

		n1DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.5gnas")
		if n1MsgContainer.N1MessageContent == nil {
			return nil, errors.New("n1MessageContainer.N1MessageContent is nil")
		}

		contentId := n1MsgContainer.N1MessageContent.ContentId
		if contentId == "" {
			return nil, errors.New("n1MessageContent.ContentId is empty")
		}

		n1DataFields[0].Set(reflect.ValueOf(&multipart.RelatedContent{
			ContentID: contentId,
			Content:   c.BinaryDataN1Message,
		}))
	}

	return m, nil
}
