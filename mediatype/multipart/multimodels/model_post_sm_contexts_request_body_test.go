package multimodels

import (
	"testing"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestPostSmContextsRequestBody_ReadFromModels(t *testing.T) {
	m := models.PostSmContextsRequestBody{}
	m.JsonData = &models.Smf_PDUSess_SmContextCreateData{
		N1SmMsg: &models.RefToBinaryData{
			ContentId: "n1SmMsg",
		},
		N2SmInfo: &models.RefToBinaryData{
			ContentId: "n2SmInfo",
		},
		N2SmInfoExt1: &models.RefToBinaryData{
			ContentId: "n2SmInfoExt1",
		},
	}

	m.BinaryDataN1SmMessage = &multipart.RelatedContent{
		ContentID: "n1SmMsg",
		Content:   []byte("testN1"),
	}
	m.BinaryDataN2SmInformation = &multipart.RelatedContent{
		ContentID: "n2SmInfo",
		Content:   []byte("testN2"),
	}
	m.BinaryDataN2SmInformationExt1 = &multipart.RelatedContent{
		ContentID: "n2SmInfoExt1",
		Content:   []byte("testN2Ext1"),
	}

	target := &PostSmContextsRequestBody{}
	err := target.ReadFromModels(&m)
	require.NoError(t, err)
	require.Equal(t, target.BinaryDataN1SmMessage, []byte("testN1"))
	require.Equal(t, target.BinaryDataN2SmInformation, []byte("testN2"))
	require.Equal(t, target.BinaryDataN2SmInformationExt1, []byte("testN2Ext1"))
}

func TestPostSmContextsRequestBody_BuildToModels(t *testing.T) {
	c := &PostSmContextsRequestBody{
		JsonData: &models.Smf_PDUSess_SmContextCreateData{
			N1SmMsg: &models.RefToBinaryData{
				ContentId: "n1SmMsg",
			},
			N2SmInfo: &models.RefToBinaryData{
				ContentId: "n2SmInfo",
			},
			N2SmInfoExt1: &models.RefToBinaryData{
				ContentId: "n2SmInfoExt1",
			},
		},
	}
	c.BinaryDataN1SmMessage = []byte("testN1")
	c.BinaryDataN2SmInformation = []byte("testN2")
	c.BinaryDataN2SmInformationExt1 = []byte("testN2Ext1")

	m, err := c.BuildToModels()
	require.NoError(t, err)
	require.Equal(t, m.JsonData.N1SmMsg.ContentId, "n1SmMsg")
	require.Equal(t, m.JsonData.N2SmInfo.ContentId, "n2SmInfo")
	require.Equal(t, m.JsonData.N2SmInfoExt1.ContentId, "n2SmInfoExt1")

	require.Equal(t, m.BinaryDataN1SmMessage.ContentID, "n1SmMsg")
	require.Equal(t, m.BinaryDataN1SmMessage.Content, []byte("testN1"))

	require.Equal(t, m.BinaryDataN2SmInformation.ContentID, "n2SmInfo")
	require.Equal(t, m.BinaryDataN2SmInformation.Content, []byte("testN2"))

	require.Equal(t, m.BinaryDataN2SmInformationExt1.ContentID, "n2SmInfoExt1")
	require.Equal(t, m.BinaryDataN2SmInformationExt1.Content, []byte("testN2Ext1"))
}
