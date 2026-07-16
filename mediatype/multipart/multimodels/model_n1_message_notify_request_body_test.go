package multimodels

import (
	"testing"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestN1MessageNotifyRequestBody_ReadFromModels(t *testing.T) {
	m := models.N1MessageNotifyRequestBody{}
	m.JsonData = &models.Amf_Comm_N1MessageNotification{
		N1MessageContainer: &models.Amf_Comm_N1MessageContainer{
			N1MessageContent: &models.RefToBinaryData{ContentId: "5GNAS"},
		},
	}
	m.BinaryDataN1Message = &multipart.RelatedContent{
		ContentID: "5GNAS",
		Content:   []byte("5GNAS"),
	}

	c := &N1MessageNotifyRequestBody{}
	err := c.ReadFromModels(&m)
	require.NoError(t, err)

	require.Equal(t, c.BinaryDataN1Message, []byte("5GNAS"))
}

func TestN1MessageNotifyRequestBody_BuildToModels(t *testing.T) {
	c := &N1MessageNotifyRequestBody{}
	c.JsonData = &models.Amf_Comm_N1MessageNotification{
		N1MessageContainer: &models.Amf_Comm_N1MessageContainer{
			N1MessageContent: &models.RefToBinaryData{ContentId: "5GNAS"},
		},
	}
	c.BinaryDataN1Message = []byte("5GNAS")

	m, err := c.BuildToModels()
	require.NoError(t, err)

	require.Equal(t, m.JsonData.N1MessageContainer.N1MessageContent.ContentId, "5GNAS")
	require.Equal(t, m.BinaryDataN1Message.ContentID, "5GNAS")
	require.Equal(t, m.BinaryDataN1Message.Content, []byte("5GNAS"))
}
