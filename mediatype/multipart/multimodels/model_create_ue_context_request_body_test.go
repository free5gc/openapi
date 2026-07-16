package multimodels

import (
	"testing"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestCreateUEContextRequestBody_ReadFromModels(t *testing.T) {
	m := models.CreateUEContextRequestBody{}
	m.JsonData = &models.Amf_Comm_UeContextCreateData{
		PduSessionList: []models.Amf_Comm_N2SmInformation{
			{
				PduSessionId: 1,
				N2InfoContent: &models.Amf_Comm_N2InfoContent{
					NgapData: &models.RefToBinaryData{
						ContentId: "PDU_SESSION_1",
					},
				},
			},
			{
				PduSessionId: 2,
				N2InfoContent: &models.Amf_Comm_N2InfoContent{
					NgapData: &models.RefToBinaryData{
						ContentId: "PDU_SESSION_2",
					},
				},
			},
			{
				PduSessionId: 10,
				N2InfoContent: &models.Amf_Comm_N2InfoContent{
					NgapData: &models.RefToBinaryData{
						ContentId: "PDU_SESSION_10",
					},
				},
			},
		},
		UeRadioCapability: &models.Amf_Comm_N2InfoContent{
			NgapData: &models.RefToBinaryData{
				ContentId: "UE_RADIO_CAPABILITY",
			},
		},
	}
	m.BinaryDataN2Information = &multipart.RelatedContent{
		ContentID: "PDU_SESSION_1",
		Content:   []byte("test1"),
	}
	m.BinaryDataN2InformationExt1 = &multipart.RelatedContent{
		ContentID: "PDU_SESSION_2",
		Content:   []byte("test2"),
	}
	m.BinaryDataN2InformationExt10 = &multipart.RelatedContent{
		ContentID: "PDU_SESSION_10",
		Content:   []byte("test10"),
	}

	m.BinaryDataN2InformationExt17 = &multipart.RelatedContent{
		ContentID: "UE_RADIO_CAPABILITY",
		Content:   []byte("ueRadioCapability"),
	}

	target := &CreateUEContextRequestBody{}
	err := target.ReadFromModels(&m)
	require.NoError(t, err)
	require.Equal(t, target.PduSessionN2Information[1], []byte("test1"))
	require.Equal(t, target.PduSessionN2Information[2], []byte("test2"))
	require.Equal(t, target.PduSessionN2Information[10], []byte("test10"))
	require.Equal(t, target.UeRadioCapabilityN2Information, []byte("ueRadioCapability"))
}

func TestCreateUEContextRequestBody_BuildToModels(t *testing.T) {
	c := &CreateUEContextRequestBody{
		JsonData: &models.Amf_Comm_UeContextCreateData{
			PduSessionList: []models.Amf_Comm_N2SmInformation{
				{
					PduSessionId: 1,
					N2InfoContent: &models.Amf_Comm_N2InfoContent{
						NgapData: &models.RefToBinaryData{
							ContentId: "PDU_SESSION_1",
						},
					},
				},
				{
					PduSessionId: 10,
					N2InfoContent: &models.Amf_Comm_N2InfoContent{
						NgapData: &models.RefToBinaryData{
							ContentId: "PDU_SESSION_10",
						},
					},
				},
			},
			UeRadioCapability: &models.Amf_Comm_N2InfoContent{
				NgapData: &models.RefToBinaryData{
					ContentId: "UE_RADIO_CAPABILITY",
				},
			},
		},
	}
	c.PduSessionN2Information[1] = []byte("test1")
	c.PduSessionN2Information[10] = []byte("test10")
	c.UeRadioCapabilityN2Information = []byte("ueRadioCapability")

	m, err := c.BuildToModels()
	require.NoError(t, err)
	require.Equal(t, m.JsonData.PduSessionList[0].N2InfoContent.NgapData.ContentId, "PDU_SESSION_1")
	require.Equal(t, m.JsonData.PduSessionList[1].N2InfoContent.NgapData.ContentId, "PDU_SESSION_10")
	require.Equal(t, m.JsonData.UeRadioCapability.NgapData.ContentId, "UE_RADIO_CAPABILITY")

	require.Equal(t, m.BinaryDataN2Information.ContentID, "PDU_SESSION_1")
	require.Equal(t, m.BinaryDataN2Information.Content, []byte("test1"))

	require.Equal(t, m.BinaryDataN2InformationExt1.ContentID, "PDU_SESSION_10")
	require.Equal(t, m.BinaryDataN2InformationExt1.Content, []byte("test10"))

	require.Equal(t, m.BinaryDataN2InformationExt2.ContentID, "UE_RADIO_CAPABILITY")
	require.Equal(t, m.BinaryDataN2InformationExt2.Content, []byte("ueRadioCapability"))
}
