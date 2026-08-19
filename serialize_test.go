package openapi

import (
	"mime"
	"testing"

	"github.com/free5gc/openapi/mediatype/multipart"
	"github.com/stretchr/testify/require"
)

type testJSON struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type testMultipartRelated struct {
	JsonData   *testJSON                 `json:"jsonData" multipart:"contentType:application/json"`
	NgapBinary *multipart.RelatedContent `json:"ngapBinary" multipart:"contentType:application/vnd.3gpp.ngap"`
}

func TestSerialize_JSON(t *testing.T) {
	body := testJSON{
		Name: "John",
		Age:  30,
	}
	contentType, bodyBytes, err := Serialize(body, "application/json")
	require.NoError(t, err)
	require.Equal(t, "application/json", contentType)
	require.Equal(t, `{"name":"John","age":30}`, string(bodyBytes))
}

func TestSerialize_MultipartRelated(t *testing.T) {
	body := testMultipartRelated{
		JsonData: &testJSON{
			Name: "John",
			Age:  30,
		},
		NgapBinary: &multipart.RelatedContent{
			ContentID: "ngapBinary",
			Content:   []byte("ngapBinary"),
		},
	}
	contentType, bodyBytes, err := Serialize(body, "multipart/related")
	require.NoError(t, err)
	mediaType, params, err := mime.ParseMediaType(contentType)
	require.NoError(t, err)
	require.Equal(t, "multipart/related", mediaType)
	require.Contains(t, params, "boundary")
	require.Contains(t, string(bodyBytes), `{"name":"John","age":30}`)
	require.Contains(t, string(bodyBytes), `application/vnd.3gpp.ngap`)
}
