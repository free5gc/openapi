package multimodels

import (
	"testing"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestDetermineLocationRequestBody_ReadFromModels(t *testing.T) {
	m := models.DetermineLocationRequestBody{}
	m.JsonData = &models.Lmf_Loc_InputData{
		LppMessage: &models.RefToBinaryData{
			ContentId: "lppMessage",
		},
		LppMessageExt: []models.RefToBinaryData{
			{
				ContentId: "lppMessageExt1",
			},
			{
				ContentId: "lppMessageExt2",
			},
		},
	}

	m.BinaryDataLppMessage = &multipart.RelatedContent{
		ContentID: "lppMessage",
		Content:   []byte("testLpp"),
	}
	m.BinaryDataLppMessageExt1 = &multipart.RelatedContent{
		ContentID: "lppMessageExt1",
		Content:   []byte("testLppExt1"),
	}
	m.BinaryDataLppMessageExt2 = &multipart.RelatedContent{
		ContentID: "lppMessageExt2",
		Content:   []byte("testLppExt2"),
	}

	target := &DetermineLocationRequestBody{}
	err := target.ReadFromModels(&m)
	require.NoError(t, err)
	require.Equal(t, target.BinaryDataLppMessage, []byte("testLpp"))
	require.Equal(t, target.BinaryDataLppMessageExt1, []byte("testLppExt1"))
	require.Equal(t, target.BinaryDataLppMessageExt2, []byte("testLppExt2"))
}

func TestDetermineLocationRequestBody_ReadFromModels_SingleLppMessage(t *testing.T) {
	m := models.DetermineLocationRequestBody{}
	m.JsonData = &models.Lmf_Loc_InputData{
		LppMessage: &models.RefToBinaryData{
			ContentId: "lppMessage",
		},
	}

	m.BinaryDataLppMessage = &multipart.RelatedContent{
		ContentID: "lppMessage",
		Content:   []byte("testLpp"),
	}

	target := &DetermineLocationRequestBody{}
	err := target.ReadFromModels(&m)
	require.NoError(t, err)
	require.Equal(t, target.BinaryDataLppMessage, []byte("testLpp"))
	require.Nil(t, target.BinaryDataLppMessageExt1)
	require.Nil(t, target.BinaryDataLppMessageExt2)
}

func TestDetermineLocationRequestBody_BuildToModels(t *testing.T) {
	c := &DetermineLocationRequestBody{
		JsonData: &models.Lmf_Loc_InputData{
			LppMessage: &models.RefToBinaryData{
				ContentId: "lppMessage",
			},
			LppMessageExt: []models.RefToBinaryData{
				{
					ContentId: "lppMessageExt1",
				},
				{
					ContentId: "lppMessageExt2",
				},
			},
		},
	}
	c.BinaryDataLppMessage = []byte("testLpp")
	c.BinaryDataLppMessageExt1 = []byte("testLppExt1")
	c.BinaryDataLppMessageExt2 = []byte("testLppExt2")

	m, err := c.BuildToModels()
	require.NoError(t, err)
	require.Equal(t, m.JsonData.LppMessage.ContentId, "lppMessage")
	require.Equal(t, m.JsonData.LppMessageExt[0].ContentId, "lppMessageExt1")
	require.Equal(t, m.JsonData.LppMessageExt[1].ContentId, "lppMessageExt2")

	require.Equal(t, m.BinaryDataLppMessage.ContentID, "lppMessage")
	require.Equal(t, m.BinaryDataLppMessage.Content, []byte("testLpp"))

	require.Equal(t, m.BinaryDataLppMessageExt1.ContentID, "lppMessageExt1")
	require.Equal(t, m.BinaryDataLppMessageExt1.Content, []byte("testLppExt1"))

	require.Equal(t, m.BinaryDataLppMessageExt2.ContentID, "lppMessageExt2")
	require.Equal(t, m.BinaryDataLppMessageExt2.Content, []byte("testLppExt2"))
}

func TestDetermineLocationRequestBody_BuildToModels_SingleLppMessage(t *testing.T) {
	c := &DetermineLocationRequestBody{
		JsonData: &models.Lmf_Loc_InputData{
			LppMessage: &models.RefToBinaryData{
				ContentId: "lppMessage",
			},
		},
	}
	c.BinaryDataLppMessage = []byte("testLpp")

	m, err := c.BuildToModels()
	require.NoError(t, err)
	require.Equal(t, m.JsonData.LppMessage.ContentId, "lppMessage")

	require.Equal(t, m.BinaryDataLppMessage.ContentID, "lppMessage")
	require.Equal(t, m.BinaryDataLppMessage.Content, []byte("testLpp"))

	require.Nil(t, m.BinaryDataLppMessageExt1)
	require.Nil(t, m.BinaryDataLppMessageExt2)
}
