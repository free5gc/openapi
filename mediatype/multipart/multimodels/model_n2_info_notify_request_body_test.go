package multimodels

import (
	"testing"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestN2InfoNotifyRequestBody_ReadFromModels(t *testing.T) {
	m := &models.N2InfoNotifyRequestBody{}
	m.JsonData = &models.Amf_Comm_N2InformationNotification{
		N2InfoContainer: &models.Amf_Comm_N2InfoContainer{
			N2InformationClass: models.Amf_Comm_N2InformationClass_SM,
			SmInfo: &models.Amf_Comm_N2SmInformation{
				N2InfoContent: &models.Amf_Comm_N2InfoContent{
					NgapData: &models.RefToBinaryData{ContentId: "SM-NGAP"},
				},
			},
		},
	}
	m.BinaryDataN2Information = &multipart.RelatedContent{
		ContentID: "SM-NGAP",
		Content:   []byte("SM-NGAP"),
	}

	c := &N2InfoNotifyRequestBody{}
	err := c.ReadFromModels(m)
	require.NoError(t, err)

	require.Equal(t, []byte("SM-NGAP"), c.BinaryDataN2Information)
}

func TestN2InfoNotifyRequestBody_BuildToModels(t *testing.T) {
	c := &N2InfoNotifyRequestBody{}
	c.JsonData = &models.Amf_Comm_N2InformationNotification{
		N2InfoContainer: &models.Amf_Comm_N2InfoContainer{
			N2InformationClass: models.Amf_Comm_N2InformationClass_SM,
			SmInfo: &models.Amf_Comm_N2SmInformation{
				N2InfoContent: &models.Amf_Comm_N2InfoContent{
					NgapData: &models.RefToBinaryData{ContentId: "SM-NGAP"},
				},
			},
		},
	}
	c.BinaryDataN2Information = []byte("SM-NGAP")

	m, err := c.BuildToModels()
	require.NoError(t, err)

	require.Equal(t, &multipart.RelatedContent{
		ContentID: "SM-NGAP",
		Content:   []byte("SM-NGAP"),
	}, m.BinaryDataN2Information)
}
