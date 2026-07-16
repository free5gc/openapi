package multimodels

import (
	"testing"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestUEContextTransferResponse200_ReadFromModels(t *testing.T) {
	m := &models.UEContextTransferResponse200{}
	m.JsonData = &models.Amf_Comm_UeContextTransferRspData{
		UeRadioCapability: &models.Amf_Comm_N2InfoContent{
			NgapData: &models.RefToBinaryData{
				ContentId: "UE-RADIO-CAPABILITY",
			},
		},
		UeRadioCapabilityForPaging: &models.Amf_Comm_N2InfoContent{
			NgapData: &models.RefToBinaryData{
				ContentId: "UE-RADIO-CAPABILITY-FOR-PAGING",
			},
		},
		UeNbiotRadioCapability: &models.Amf_Comm_N2InfoContent{
			NgapData: &models.RefToBinaryData{
				ContentId: "UE-NBIOT-RADIO-CAPABILITY",
			},
		},
	}

	m.BinaryDataN2Information = &multipart.RelatedContent{
		ContentID: "UE-RADIO-CAPABILITY",
		Content:   []byte("UE-RADIO-CAPABILITY"),
	}
	m.BinaryDataN2InformationExt1 = &multipart.RelatedContent{
		ContentID: "UE-RADIO-CAPABILITY-FOR-PAGING",
		Content:   []byte("UE-RADIO-CAPABILITY-FOR-PAGING"),
	}
	m.BinaryDataN2InformationExt2 = &multipart.RelatedContent{
		ContentID: "UE-NBIOT-RADIO-CAPABILITY",
		Content:   []byte("UE-NBIOT-RADIO-CAPABILITY"),
	}

	c := &UEContextTransferResponse200{}
	err := c.ReadFromModels(m)
	require.NoError(t, err)

	require.Equal(t, c.UeRadioCapabilityN2Information, []byte("UE-RADIO-CAPABILITY"))
	require.Equal(t, c.UeRadioCapabilityForPagingN2Information, []byte("UE-RADIO-CAPABILITY-FOR-PAGING"))
	require.Equal(t, c.UeNbiotRadioCapabilityN2Information, []byte("UE-NBIOT-RADIO-CAPABILITY"))
}

func TestUEContextTransferResponse200_BuildToModels(t *testing.T) {
	c := &UEContextTransferResponse200{}
	c.JsonData = &models.Amf_Comm_UeContextTransferRspData{
		UeRadioCapability: &models.Amf_Comm_N2InfoContent{
			NgapData: &models.RefToBinaryData{
				ContentId: "UE-RADIO-CAPABILITY",
			},
		},
		UeRadioCapabilityForPaging: &models.Amf_Comm_N2InfoContent{
			NgapData: &models.RefToBinaryData{
				ContentId: "UE-RADIO-CAPABILITY-FOR-PAGING",
			},
		},
		UeNbiotRadioCapability: &models.Amf_Comm_N2InfoContent{
			NgapData: &models.RefToBinaryData{
				ContentId: "UE-NBIOT-RADIO-CAPABILITY",
			},
		},
	}
	c.UeRadioCapabilityN2Information = []byte("UE-RADIO-CAPABILITY")
	c.UeRadioCapabilityForPagingN2Information = []byte("UE-RADIO-CAPABILITY-FOR-PAGING")
	c.UeNbiotRadioCapabilityN2Information = []byte("UE-NBIOT-RADIO-CAPABILITY")

	m, err := c.BuildToModels()
	require.NoError(t, err)

	require.Equal(t, m.JsonData.UeRadioCapability.NgapData.ContentId, "UE-RADIO-CAPABILITY")
	require.Equal(t, m.BinaryDataN2Information.ContentID, "UE-RADIO-CAPABILITY")
	require.Equal(t, m.BinaryDataN2Information.Content, []byte("UE-RADIO-CAPABILITY"))

	require.Equal(t, m.JsonData.UeRadioCapabilityForPaging.NgapData.ContentId, "UE-RADIO-CAPABILITY-FOR-PAGING")
	require.Equal(t, m.BinaryDataN2InformationExt1.ContentID, "UE-RADIO-CAPABILITY-FOR-PAGING")
	require.Equal(t, m.BinaryDataN2InformationExt1.Content, []byte("UE-RADIO-CAPABILITY-FOR-PAGING"))

	require.Equal(t, m.JsonData.UeNbiotRadioCapability.NgapData.ContentId, "UE-NBIOT-RADIO-CAPABILITY")
	require.Equal(t, m.BinaryDataN2InformationExt2.ContentID, "UE-NBIOT-RADIO-CAPABILITY")
	require.Equal(t, m.BinaryDataN2InformationExt2.Content, []byte("UE-NBIOT-RADIO-CAPABILITY"))
}
