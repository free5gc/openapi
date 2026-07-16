package multimodels

import (
	"testing"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestUEContextTransferRequestBody_ReadFromModels(t *testing.T) {
	m := &models.UEContextTransferRequestBody{}
	m.JsonData = &models.Amf_Comm_UeContextTransferReqData{
		RegRequest: &models.Amf_Comm_N1MessageContainer{
			N1MessageContent: &models.RefToBinaryData{ContentId: "REGISTRATION-REQUEST"},
		},
	}
	m.BinaryDataN1Message = &multipart.RelatedContent{
		ContentID: "REGISTRATION-REQUEST",
		Content:   []byte("REGISTRATION-REQUEST"),
	}

	c := &UEContextTransferRequestBody{}
	err := c.ReadFromModels(m)
	require.NoError(t, err)

	require.Equal(t, c.BinaryDataN1Message, []byte("REGISTRATION-REQUEST"))
}

func TestUEContextTransferRequestBody_BuildToModels(t *testing.T) {
	c := &UEContextTransferRequestBody{}
	c.JsonData = &models.Amf_Comm_UeContextTransferReqData{
		RegRequest: &models.Amf_Comm_N1MessageContainer{
			N1MessageContent: &models.RefToBinaryData{ContentId: "REGISTRATION-REQUEST"},
		},
	}
	c.BinaryDataN1Message = []byte("REGISTRATION-REQUEST")

	m, err := c.BuildToModels()
	require.NoError(t, err)

	require.Equal(t, m.JsonData.RegRequest.N1MessageContent.ContentId, "REGISTRATION-REQUEST")
	require.Equal(t, m.BinaryDataN1Message.ContentID, "REGISTRATION-REQUEST")
	require.Equal(t, m.BinaryDataN1Message.Content, []byte("REGISTRATION-REQUEST"))
}
