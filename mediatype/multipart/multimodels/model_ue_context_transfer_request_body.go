package multimodels

import (
	"errors"
	"reflect"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
)

type UEContextTransferRequestBody struct {
	JsonData            *models.Amf_Comm_UeContextTransferReqData `json:"jsonData,omitempty"            multipart:"contentType:application/json"           yaml:"jsonData"            bson:"jsonData,omitempty"`
	BinaryDataN1Message []byte                                    `json:"binaryDataN1Message,omitempty" multipart:"contentType:application/vnd.3gpp.5gnas" yaml:"binaryDataN1Message" bson:"binaryDataN1Message,omitempty"`
}

func (c *UEContextTransferRequestBody) ReadFromModels(m *models.UEContextTransferRequestBody) error {
	if m == nil {
		return errors.New("model is nil")
	}

	if m.JsonData == nil {
		return errors.New("jsonData is nil")
	}

	c.JsonData = m.JsonData

	regRequest := c.JsonData.RegRequest
	if regRequest == nil {
		return nil
	}

	n1MsgContent := regRequest.N1MessageContent
	if n1MsgContent == nil {
		return errors.New("n1MessageContent is nil")
	}

	n1DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.5gnas")
	c.BinaryDataN1Message = multipart.MatchRelatedContent(n1MsgContent.ContentId, n1DataFields)

	return nil
}

func (c *UEContextTransferRequestBody) BuildToModels() (*models.UEContextTransferRequestBody, error) {
	if c == nil {
		return nil, errors.New("multipart model is nil")
	}

	if c.JsonData == nil {
		return nil, errors.New("jsonData is nil")
	}

	m := &models.UEContextTransferRequestBody{}
	m.JsonData = c.JsonData

	n1DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.5gnas")

	regRequest := c.JsonData.RegRequest
	if regRequest == nil {
		return m, nil
	}

	n1MsgContent := regRequest.N1MessageContent
	if n1MsgContent == nil {
		return nil, errors.New("n1MessageContent is nil")
	}

	contentId := n1MsgContent.ContentId
	if contentId == "" {
		return nil, errors.New("n1MessageContent.ContentId is empty")
	}

	n1DataFields[0].Set(reflect.ValueOf(&multipart.RelatedContent{
		ContentID: contentId,
		Content:   c.BinaryDataN1Message,
	}))

	return m, nil
}
