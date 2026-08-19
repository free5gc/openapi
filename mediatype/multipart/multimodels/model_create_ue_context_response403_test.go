package multimodels

import (
	"testing"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestCreateUEContextResponse403_ReadFromModels(t *testing.T) {
	m := models.CreateUEContextResponse403{}
	m.JsonData = &models.Amf_Comm_UeContextCreateError{
		TargetToSourceFailureData: &models.Amf_Comm_N2InfoContent{
			NgapData: &models.RefToBinaryData{
				ContentId: "TARGET-TO-SOURCE-FAILURE-DATA",
			},
		},
	}
	m.BinaryDataN2Information = &multipart.RelatedContent{
		ContentID: "TARGET-TO-SOURCE-FAILURE-DATA",
		Content:   []byte("TARGET-TO-SOURCE-FAILURE-DATA"),
	}

	c := &CreateUEContextResponse403{}
	err := c.ReadFromModels(&m)
	require.NoError(t, err)

	require.Equal(t, c.TargetToSourceFailureN2Information, []byte("TARGET-TO-SOURCE-FAILURE-DATA"))
	require.Equal(t, c.JsonData.TargetToSourceFailureData.NgapData.ContentId, "TARGET-TO-SOURCE-FAILURE-DATA")
}

func TestCreateUEContextResponse403_BuildToModels(t *testing.T) {
	c := &CreateUEContextResponse403{}
	c.JsonData = &models.Amf_Comm_UeContextCreateError{
		TargetToSourceFailureData: &models.Amf_Comm_N2InfoContent{
			NgapData: &models.RefToBinaryData{
				ContentId: "TARGET-TO-SOURCE-FAILURE-DATA",
			},
		},
	}
	c.TargetToSourceFailureN2Information = []byte("TARGET-TO-SOURCE-FAILURE-DATA")

	m, err := c.BuildToModels()
	require.NoError(t, err)

	require.Equal(t, m.JsonData.TargetToSourceFailureData.NgapData.ContentId, "TARGET-TO-SOURCE-FAILURE-DATA")
	require.Equal(t, m.BinaryDataN2Information.ContentID, "TARGET-TO-SOURCE-FAILURE-DATA")
	require.Equal(t, m.BinaryDataN2Information.Content, []byte("TARGET-TO-SOURCE-FAILURE-DATA"))
}
