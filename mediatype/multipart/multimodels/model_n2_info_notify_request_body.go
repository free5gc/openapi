package multimodels

import (
	"errors"
	"fmt"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
)

type N2InfoNotifyRequestBody struct {
	JsonData                *models.Amf_Comm_N2InformationNotification `json:"jsonData,omitempty"                multipart:"contentType:application/json"           yaml:"jsonData"                bson:"jsonData,omitempty"`
	BinaryDataN1Message     []byte                                     `json:"binaryDataN1Message,omitempty"     multipart:"contentType:application/vnd.3gpp.5gnas" yaml:"binaryDataN1Message"     bson:"binaryDataN1Message,omitempty"`
	BinaryDataN2Information []byte                                     `json:"binaryDataN2Information,omitempty" multipart:"contentType:application/vnd.3gpp.ngap"  yaml:"binaryDataN2Information" bson:"binaryDataN2Information,omitempty"`
}

func (c *N2InfoNotifyRequestBody) ReadFromModels(m *models.N2InfoNotifyRequestBody) error {
	if m == nil {
		return errors.New("model is nil")
	}

	if m.JsonData == nil {
		return errors.New("jsonData is nil")
	}

	c.JsonData = m.JsonData

	n2DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.ngap")

	if c.JsonData.N2InfoContainer != nil {
		var err error
		c.BinaryDataN2Information, err = readN2InformationFromRelatedContent(
			c.JsonData.N2InfoContainer,
			n2DataFields,
		)
		if err != nil {
			return fmt.Errorf("failed to read N2Information: %w", err)
		}
	}
	return nil
}

func (c *N2InfoNotifyRequestBody) BuildToModels() (*models.N2InfoNotifyRequestBody, error) {
	if c == nil {
		return nil, errors.New("multipart model is nil")
	}

	if c.JsonData == nil {
		return nil, errors.New("jsonData is nil")
	}

	m := &models.N2InfoNotifyRequestBody{}
	m.JsonData = c.JsonData

	n2DataFields := multipart.GetRelatedContentFields(m, "application/vnd.3gpp.ngap")

	if c.JsonData.N2InfoContainer != nil {
		err := putN2InformationToRelatedContent(c.JsonData.N2InfoContainer, c.BinaryDataN2Information, n2DataFields[0])
		if err != nil {
			return nil, fmt.Errorf("failed to put N2Information to related content: %w", err)
		}
	}

	return m, nil
}
