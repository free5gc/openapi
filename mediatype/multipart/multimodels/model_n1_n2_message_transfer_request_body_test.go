package multimodels

import (
	"testing"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestN1N2MessageTransferRequestBody_ReadFromModels(t *testing.T) {
	m := models.N1N2MessageTransferRequestBody{}
	m.JsonData = &models.Amf_Comm_N1N2MessageTransferReqData{
		N1MessageContainer: &models.Amf_Comm_N1MessageContainer{
			N1MessageContent: &models.RefToBinaryData{
				ContentId: "5GNAS",
			},
		},
		N2InfoContainer: &models.Amf_Comm_N2InfoContainer{
			N2InformationClass: models.Amf_Comm_N2InformationClass_SM,
			SmInfo: &models.Amf_Comm_N2SmInformation{
				N2InfoContent: &models.Amf_Comm_N2InfoContent{
					NgapData: &models.RefToBinaryData{ContentId: "SM-NGAP"},
				},
			},
		},
		MtData: &models.RefToBinaryData{ContentId: "MT-5GNAS"},
	}
	m.BinaryDataN1Message = &multipart.RelatedContent{
		ContentID: "5GNAS",
		Content:   []byte("5GNAS"),
	}
	m.BinaryDataN2Information = &multipart.RelatedContent{
		ContentID: "SM-NGAP",
		Content:   []byte("SM-NGAP"),
	}
	m.BinaryMtData = &multipart.RelatedContent{
		ContentID: "MT-5GNAS",
		Content:   []byte("MT-5GNAS"),
	}

	c := &N1N2MessageTransferRequestBody{}
	err := c.ReadFromModels(&m)
	require.NoError(t, err)

	require.Equal(t, []byte("5GNAS"), c.N1Message)
	require.Equal(t, []byte("SM-NGAP"), c.N2Information)
	require.Equal(t, []byte("MT-5GNAS"), c.MtData)
}

func TestN1N2MessageTransferRequestBody_BuildToModels(t *testing.T) {
	c := &N1N2MessageTransferRequestBody{}
	c.JsonData = &models.Amf_Comm_N1N2MessageTransferReqData{
		N1MessageContainer: &models.Amf_Comm_N1MessageContainer{
			N1MessageContent: &models.RefToBinaryData{ContentId: "5GNAS"},
		},
		N2InfoContainer: &models.Amf_Comm_N2InfoContainer{
			N2InformationClass: models.Amf_Comm_N2InformationClass_SM,
			SmInfo: &models.Amf_Comm_N2SmInformation{
				N2InfoContent: &models.Amf_Comm_N2InfoContent{
					NgapData: &models.RefToBinaryData{ContentId: "SM-NGAP"},
				},
			},
		},
		MtData: &models.RefToBinaryData{ContentId: "MT-5GNAS"},
	}
	c.N1Message = []byte("5GNAS")
	c.N2Information = []byte("SM-NGAP")
	c.MtData = []byte("MT-5GNAS")

	m, err := c.BuildToModels()
	require.NoError(t, err)

	require.Equal(t, m.JsonData.N1MessageContainer.N1MessageContent.ContentId, "5GNAS")
	require.Equal(t, m.JsonData.N2InfoContainer.N2InformationClass, models.Amf_Comm_N2InformationClass_SM)
	require.Equal(t, m.JsonData.N2InfoContainer.SmInfo.N2InfoContent.NgapData.ContentId, "SM-NGAP")
	require.Equal(t, m.JsonData.MtData.ContentId, "MT-5GNAS")
}
