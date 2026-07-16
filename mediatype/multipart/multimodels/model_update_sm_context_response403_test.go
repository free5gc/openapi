package multimodels

import (
	"testing"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestUpdateSmContextResponse403_ReadFromModels(t *testing.T) {
	m := models.UpdateSmContextResponse403{}
	m.JsonData = &models.Smf_PDUSess_SmContextUpdateError{
		N1SmMsg: &models.RefToBinaryData{
			ContentId: "n1SmMsg",
		},
		N2SmInfo: &models.RefToBinaryData{
			ContentId: "n2SmInfo",
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

	target := &UpdateSmContextResponse403{}
	err := target.ReadFromModels(&m)
	require.NoError(t, err)
	require.Equal(t, target.BinaryDataN1SmMessage, []byte("testN1"))
	require.Equal(t, target.BinaryDataN2SmInformation, []byte("testN2"))
}

func TestUpdateSmContextResponse403_BuildToModels(t *testing.T) {
	c := &UpdateSmContextResponse403{
		JsonData: &models.Smf_PDUSess_SmContextUpdateError{
			N1SmMsg: &models.RefToBinaryData{
				ContentId: "n1SmMsg",
			},
			N2SmInfo: &models.RefToBinaryData{
				ContentId: "n2SmInfo",
			},
		},
	}
	c.BinaryDataN1SmMessage = []byte("testN1")
	c.BinaryDataN2SmInformation = []byte("testN2")

	m, err := c.BuildToModels()
	require.NoError(t, err)
	require.Equal(t, m.JsonData.N2SmInfo.ContentId, "n2SmInfo")
	require.Equal(t, m.BinaryDataN1SmMessage.ContentID, "n1SmMsg")
	require.Equal(t, m.BinaryDataN2SmInformation.ContentID, "n2SmInfo")
	require.Equal(t, m.BinaryDataN1SmMessage.Content, []byte("testN1"))
	require.Equal(t, m.BinaryDataN2SmInformation.Content, []byte("testN2"))
}
