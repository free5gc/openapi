package multimodels

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
)

type CreateUEContextResponse201 struct {
	JsonData                      *models.Amf_Comm_UeContextCreatedData `json:"jsonData,omitempty"                     multipart:"contentType:application/json"          yaml:"jsonData"                     bson:"jsonData,omitempty"`
	TargetToSourceN2Information   []byte                                `json:"targetToSourceData,omitempty"           multipart:"contentType:application/vnd.3gpp.ngap" yaml:"targetToSourceData"           bson:"targetToSourceData,omitempty"`
	PduSessionN2Information       [16][]byte                            `json:"pduSessionN2Information,omitempty"      multipart:"contentType:application/vnd.3gpp.ngap" yaml:"pduSessionN2Information"      bson:"pduSessionN2Information,omitempty"`
	FailedPduSessionN2Information [16][]byte                            `json:"failedPduSessionN2Information,omitempty" multipart:"contentType:application/vnd.3gpp.ngap" yaml:"failedPduSessionN2Information" bson:"failedPduSessionN2Information,omitempty"`
}

func (c *CreateUEContextResponse201) ReadFromModels(m *models.CreateUEContextResponse201) error {
	if m == nil {
		return errors.New("model is nil")
	}

	if m.JsonData == nil {
		return errors.New("jsonData is nil")
	}

	c.JsonData = m.JsonData

	n2DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.ngap")

	if c.JsonData.TargetToSourceData != nil {
		contentId := c.JsonData.TargetToSourceData.NgapData.ContentId
		if contentId == "" {
			return errors.New("targetToSourceData.NgapData.ContentId is empty")
		}
		content := multipart.MatchRelatedContent(contentId, n2DataFields)
		if content != nil {
			c.TargetToSourceN2Information = content
		}
	}

	for idx, pduSess := range c.JsonData.PduSessionList {
		pduSessId := pduSess.PduSessionId
		contentId := pduSess.N2InfoContent.NgapData.ContentId
		if contentId == "" {
			return fmt.Errorf("pduSessionList[%d].N2InfoContent.NgapData.ContentId is empty", idx)
		}
		content := multipart.MatchRelatedContent(contentId, n2DataFields)
		if content != nil {
			c.PduSessionN2Information[pduSessId] = content
		}
	}

	for idx, pduSess := range c.JsonData.FailedSessionList {
		pduSessId := pduSess.PduSessionId
		contentId := pduSess.N2InfoContent.NgapData.ContentId
		if contentId == "" {
			return fmt.Errorf("failedSessionList[%d].N2InfoContent.NgapData.ContentId is empty", idx)
		}
		content := multipart.MatchRelatedContent(contentId, n2DataFields)
		if content != nil {
			c.FailedPduSessionN2Information[pduSessId] = content
		}
	}

	return nil
}

func (c *CreateUEContextResponse201) BuildToModels() (*models.CreateUEContextResponse201, error) {
	if c == nil {
		return nil, errors.New("multipart model is nil")
	}

	if c.JsonData == nil {
		return nil, errors.New("jsonData is nil")
	}

	m := &models.CreateUEContextResponse201{}
	m.JsonData = c.JsonData

	n2DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.ngap")
	n2FieldIndex := 0

	if c.JsonData.TargetToSourceData != nil {
		if c.JsonData.TargetToSourceData.NgapData == nil {
			return nil, errors.New("targetToSourceData.NgapData is nil")
		}

		contentId := c.JsonData.TargetToSourceData.NgapData.ContentId
		if contentId == "" {
			return nil, errors.New("targetToSourceData.NgapData.ContentId is empty")
		}
		n2DataFields[n2FieldIndex].Set(reflect.ValueOf(&multipart.RelatedContent{
			ContentID: contentId,
			Content:   c.TargetToSourceN2Information,
		}))
		n2FieldIndex++
	}

	for idx, pduSess := range c.JsonData.PduSessionList {
		pduSessId := pduSess.PduSessionId
		if pduSess.N2InfoContent == nil {
			return nil, fmt.Errorf("pduSessionList[%d].N2InfoContent is nil", idx)
		}
		if pduSess.N2InfoContent.NgapData == nil {
			return nil, fmt.Errorf("pduSessionList[%d].N2InfoContent.NgapData is nil", idx)
		}
		contentId := pduSess.N2InfoContent.NgapData.ContentId
		if contentId == "" {
			return nil, fmt.Errorf("pduSessionList[%d].N2InfoContent.NgapData.ContentId is empty", idx)
		}
		n2DataFields[n2FieldIndex].Set(reflect.ValueOf(&multipart.RelatedContent{
			ContentID: contentId,
			Content:   c.PduSessionN2Information[pduSessId],
		}))
		n2FieldIndex++
	}

	for idx, pduSess := range c.JsonData.FailedSessionList {
		pduSessId := pduSess.PduSessionId
		if pduSess.N2InfoContent == nil {
			return nil, fmt.Errorf("failedSessionList[%d].N2InfoContent is nil", idx)
		}
		if pduSess.N2InfoContent.NgapData == nil {
			return nil, fmt.Errorf("failedSessionList[%d].N2InfoContent.NgapData is nil", idx)
		}
		contentId := pduSess.N2InfoContent.NgapData.ContentId
		if contentId == "" {
			return nil, fmt.Errorf("failedSessionList[%d].N2InfoContent.NgapData.ContentId is empty", idx)
		}
		n2DataFields[n2FieldIndex].Set(reflect.ValueOf(&multipart.RelatedContent{
			ContentID: contentId,
			Content:   c.FailedPduSessionN2Information[pduSessId],
		}))
		n2FieldIndex++
	}

	return m, nil
}
