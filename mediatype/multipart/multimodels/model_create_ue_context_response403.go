package multimodels

import (
	"errors"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
)

type CreateUEContextResponse403 struct {
	JsonData                           *models.Amf_Comm_UeContextCreateError `json:"jsonData,omitempty"                multipart:"contentType:application/json"          yaml:"jsonData"                bson:"jsonData,omitempty"`
	TargetToSourceFailureN2Information []byte                                `json:"targetToSourceFailureN2Information,omitempty" multipart:"contentType:application/vnd.3gpp.ngap" yaml:"targetToSourceFailureN2Information" bson:"targetToSourceFailureN2Information,omitempty"`
}

func (c *CreateUEContextResponse403) ReadFromModels(m *models.CreateUEContextResponse403) error {
	if m == nil {
		return errors.New("model is nil")
	}

	if m.JsonData == nil {
		return errors.New("jsonData is nil")
	}

	c.JsonData = m.JsonData

	n2DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.ngap")

	if c.JsonData.TargetToSourceFailureData != nil {
		if c.JsonData.TargetToSourceFailureData.NgapData == nil {
			return errors.New("targetToSourceFailureData.NgapData is nil")
		}
		c.TargetToSourceFailureN2Information =
			multipart.MatchRelatedContent(c.JsonData.TargetToSourceFailureData.NgapData.ContentId, n2DataFields)
	}

	return nil
}

func (c *CreateUEContextResponse403) BuildToModels() (*models.CreateUEContextResponse403, error) {
	if c == nil {
		return nil, errors.New("multipart model is nil")
	}

	if c.JsonData == nil {
		return nil, errors.New("jsonData is nil")
	}

	m := &models.CreateUEContextResponse403{}
	m.JsonData = c.JsonData

	if c.JsonData.TargetToSourceFailureData != nil {
		if c.JsonData.TargetToSourceFailureData.NgapData == nil {
			return nil, errors.New("targetToSourceFailureData.NgapData is nil")
		}
		contentId := c.JsonData.TargetToSourceFailureData.NgapData.ContentId
		if contentId == "" {
			return nil, errors.New("targetToSourceFailureData.NgapData.ContentId is empty")
		}
		m.BinaryDataN2Information = &multipart.RelatedContent{
			ContentID: contentId,
			Content:   c.TargetToSourceFailureN2Information,
		}
	}
	return m, nil
}
