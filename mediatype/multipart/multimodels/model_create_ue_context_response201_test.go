package multimodels

import (
	"testing"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestCreateUEContextResponse201_ReadFromModels(t *testing.T) {
	m := models.CreateUEContextResponse201{}
	m.JsonData = &models.Amf_Comm_UeContextCreatedData{
		TargetToSourceData: &models.Amf_Comm_N2InfoContent{
			NgapData: &models.RefToBinaryData{
				ContentId: "TARGET-TO-SOURCE-DATA",
			},
		},
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
		FailedSessionList: []models.Amf_Comm_N2SmInformation{
			{
				PduSessionId: 2,
				N2InfoContent: &models.Amf_Comm_N2InfoContent{
					NgapData: &models.RefToBinaryData{
						ContentId: "PDU_SESSION_2",
					},
				},
			},
		},
	}

	m.BinaryDataN2Information = &multipart.RelatedContent{
		ContentID: "TARGET-TO-SOURCE-DATA",
		Content:   []byte("TARGET-TO-SOURCE-DATA"),
	}
	m.BinaryDataN2InformationExt1 = &multipart.RelatedContent{
		ContentID: "PDU_SESSION_1",
		Content:   []byte("PDU_SESSION_1"),
	}
	m.BinaryDataN2InformationExt2 = &multipart.RelatedContent{
		ContentID: "PDU_SESSION_10",
		Content:   []byte("PDU_SESSION_10"),
	}
	m.BinaryDataN2InformationExt3 = &multipart.RelatedContent{
		ContentID: "PDU_SESSION_2",
		Content:   []byte("PDU_SESSION_2"),
	}

	c := &CreateUEContextResponse201{}
	err := c.ReadFromModels(&m)
	require.NoError(t, err)

	require.Equal(t, c.TargetToSourceN2Information, []byte("TARGET-TO-SOURCE-DATA"))
	require.Equal(t, c.PduSessionN2Information[1], []byte("PDU_SESSION_1"))
	require.Equal(t, c.PduSessionN2Information[10], []byte("PDU_SESSION_10"))
	require.Equal(t, c.FailedPduSessionN2Information[2], []byte("PDU_SESSION_2"))
}

func TestCreateUEContextResponse201_BuildToModels(t *testing.T) {
	c := &CreateUEContextResponse201{}
	c.JsonData = &models.Amf_Comm_UeContextCreatedData{
		TargetToSourceData: &models.Amf_Comm_N2InfoContent{
			NgapData: &models.RefToBinaryData{
				ContentId: "TARGET-TO-SOURCE-DATA",
			},
		},
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
		FailedSessionList: []models.Amf_Comm_N2SmInformation{
			{
				PduSessionId: 2,
				N2InfoContent: &models.Amf_Comm_N2InfoContent{
					NgapData: &models.RefToBinaryData{
						ContentId: "PDU_SESSION_2",
					},
				},
			},
		},
	}

	c.TargetToSourceN2Information = []byte("TARGET-TO-SOURCE-DATA")
	c.PduSessionN2Information[1] = []byte("PDU_SESSION_1")
	c.PduSessionN2Information[10] = []byte("PDU_SESSION_10")
	c.FailedPduSessionN2Information[2] = []byte("PDU_SESSION_2")

	m, err := c.BuildToModels()
	require.NoError(t, err)

	require.Equal(t, "TARGET-TO-SOURCE-DATA", m.JsonData.TargetToSourceData.NgapData.ContentId)
	require.Equal(t, "TARGET-TO-SOURCE-DATA", m.BinaryDataN2Information.ContentID)
	require.Equal(t, []byte("TARGET-TO-SOURCE-DATA"), m.BinaryDataN2Information.Content)

	require.Equal(t, "PDU_SESSION_1", m.JsonData.PduSessionList[0].N2InfoContent.NgapData.ContentId)
	require.Equal(t, "PDU_SESSION_1", m.BinaryDataN2InformationExt1.ContentID)
	require.Equal(t, []byte("PDU_SESSION_1"), m.BinaryDataN2InformationExt1.Content)

	require.Equal(t, "PDU_SESSION_10", m.JsonData.PduSessionList[1].N2InfoContent.NgapData.ContentId)
	require.Equal(t, "PDU_SESSION_10", m.BinaryDataN2InformationExt2.ContentID)
	require.Equal(t, []byte("PDU_SESSION_10"), m.BinaryDataN2InformationExt2.Content)

	require.Equal(t, "PDU_SESSION_2", m.JsonData.FailedSessionList[0].N2InfoContent.NgapData.ContentId)
	require.Equal(t, "PDU_SESSION_2", m.BinaryDataN2InformationExt3.ContentID)
	require.Equal(t, []byte("PDU_SESSION_2"), m.BinaryDataN2InformationExt3.Content)
}
