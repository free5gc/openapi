package multimodels

import (
	"testing"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestPostSmContextsResponse500_ReadFromModels(t *testing.T) {
	m := models.PostSmContextsResponse500{}
	m.JsonData = &models.Smf_PDUSess_SmContextCreateError{
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

	m.BinaryDataN2SmMessage = &multipart.RelatedContent{
		ContentID: "n2SmInfo",
		Content:   []byte("testN2"),
	}

	target := &PostSmContextsResponse500{}
	err := target.ReadFromModels(&m)
	require.NoError(t, err)
	require.Equal(t, target.BinaryDataN1SmMessage, []byte("testN1"))
	require.Equal(t, target.BinaryDataN2SmMessage, []byte("testN2"))
}

func TestPostSmContextsResponse500_BuildToModels(t *testing.T) {
	c := &PostSmContextsResponse500{
		JsonData: &models.Smf_PDUSess_SmContextCreateError{
			N1SmMsg: &models.RefToBinaryData{
				ContentId: "n1SmMsg",
			},
			N2SmInfo: &models.RefToBinaryData{
				ContentId: "n2SmInfo",
			},
		},
	}
	c.BinaryDataN1SmMessage = []byte("testN1")
	c.BinaryDataN2SmMessage = []byte("testN2")

	m, err := c.BuildToModels()
	require.NoError(t, err)
	require.Equal(t, m.JsonData.N2SmInfo.ContentId, "n2SmInfo")
	require.Equal(t, m.BinaryDataN1SmMessage.ContentID, "n1SmMsg")
	require.Equal(t, m.BinaryDataN2SmMessage.ContentID, "n2SmInfo")
	require.Equal(t, m.BinaryDataN1SmMessage.Content, []byte("testN1"))
	require.Equal(t, m.BinaryDataN2SmMessage.Content, []byte("testN2"))
}
