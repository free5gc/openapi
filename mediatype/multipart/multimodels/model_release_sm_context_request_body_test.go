package multimodels

import (
	"testing"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestReleaseSmContextRequestBody_ReadFromModels(t *testing.T) {
	m := models.ReleaseSmContextRequestBody{}
	m.JsonData = &models.Smf_PDUSess_SmContextReleaseData{
		N2SmInfo: &models.RefToBinaryData{
			ContentId: "n2SmInfo",
		},
	}

	m.BinaryDataN2SmInformation = &multipart.RelatedContent{
		ContentID: "n2SmInfo",
		Content:   []byte("testN2"),
	}

	target := &ReleaseSmContextRequestBody{}
	err := target.ReadFromModels(&m)
	require.NoError(t, err)
	require.Equal(t, target.BinaryDataN2SmInformation, []byte("testN2"))
}

func TestReleaseSmContextRequestBody_BuildToModels(t *testing.T) {
	c := &ReleaseSmContextRequestBody{
		JsonData: &models.Smf_PDUSess_SmContextReleaseData{
			N2SmInfo: &models.RefToBinaryData{
				ContentId: "n2SmInfo",
			},
		},
	}
	c.BinaryDataN2SmInformation = []byte("testN2")

	m, err := c.BuildToModels()
	require.NoError(t, err)
	require.Equal(t, m.JsonData.N2SmInfo.ContentId, "n2SmInfo")
	require.Equal(t, m.BinaryDataN2SmInformation.ContentID, "n2SmInfo")
	require.Equal(t, m.BinaryDataN2SmInformation.Content, []byte("testN2"))
}
