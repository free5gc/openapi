package multipart

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultipartRelated_Marshal(t *testing.T) {
	body := &MultipartRelatedTestBody{
		JsonData: &ContentDescriptor{
			MessageContainer:  &RefToBinaryData{ContentID: "message"},
			UserDataContainer: &RefToBinaryData{ContentID: "userData"},
		},
		Message: &RelatedContent{
			Content:   []byte("Hello, World!"),
			ContentID: "message",
		},
		UserData: &RelatedContent{
			Content:   []byte("username=test&password=1234567890"),
			ContentID: "userData",
		},
	}

	boundary, marshal, err := Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	expected := testedMultipartRelatedBytes(t, boundary)
	require.Equal(t, string(expected), string(marshal))
}

func TestMultipartRelated_Unmarshal(t *testing.T) {
	var err error
	boundary, err := generateBoundary()
	require.NoError(t, err)

	payload := testedMultipartRelatedBytes(t, boundary)

	body := &MultipartRelatedTestBody{}
	err = Unmarshal(boundary, payload, body)
	require.NoError(t, err)

	require.NotNil(t, body.JsonData)
	require.NotNil(t, body.JsonData.MessageContainer)
	require.Equal(t, body.JsonData.MessageContainer.ContentID, "message")
	require.NotNil(t, body.JsonData.UserDataContainer)
	require.Equal(t, body.JsonData.UserDataContainer.ContentID, "userData")

	require.NotNil(t, body.Message)
	require.Equal(t, body.Message.ContentID, "message")
	require.Equal(t, string(body.Message.Content), "Hello, World!")

	require.NotNil(t, body.UserData)
	require.Equal(t, body.UserData.ContentID, "userData")
	require.Equal(t, string(body.UserData.Content), "username=test&password=1234567890")
}

func TestMultipartRelated(t *testing.T) {
	body := &MultipartRelatedTestBody2{
		JsonData: &ContentDescriptor2{
			MessageContainer: []*RefToBinaryData{
				{ContentID: "message1"}, {ContentID: "message2"}, {ContentID: "message3"},
			},
			UserDataContainer: []*RefToBinaryData{
				{ContentID: "userData1"}, {ContentID: "userData2"}, {ContentID: "userData3"},
			},
		},
		Message1: &RelatedContent{
			Content:   []byte("Hello, World!"),
			ContentID: "message1",
		},
		Message2: &RelatedContent{
			Content:   []byte("Hello, World!!"),
			ContentID: "message2",
		},
		Message3: &RelatedContent{
			Content:   []byte("Hello, World!!!"),
			ContentID: "message3",
		},
		UserData1: &RelatedContent{
			Content:   []byte("username=test1&password=1234567890"),
			ContentID: "userData1",
		},
		UserData2: &RelatedContent{
			Content:   []byte("username=test2&password=1234567890"),
			ContentID: "userData2",
		},
		UserData3: &RelatedContent{
			Content:   []byte("username=test3&password=1234567890"),
			ContentID: "userData3",
		},
	}

	boundary, marshal, err := Marshal(body)
	require.NoError(t, err)

	var retBody MultipartRelatedTestBody2
	err = Unmarshal(boundary, marshal, &retBody)
	require.NoError(t, err)

	require.Equal(t, retBody.JsonData.MessageContainer[0].ContentID, "message1")
	require.Equal(t, retBody.JsonData.MessageContainer[1].ContentID, "message2")
	require.Equal(t, retBody.JsonData.MessageContainer[2].ContentID, "message3")
	require.Equal(t, retBody.JsonData.UserDataContainer[0].ContentID, "userData1")
	require.Equal(t, retBody.JsonData.UserDataContainer[1].ContentID, "userData2")
	require.Equal(t, retBody.JsonData.UserDataContainer[2].ContentID, "userData3")
	require.Equal(t, retBody.Message1.ContentID, "message1")
	require.Equal(t, retBody.Message2.ContentID, "message2")
	require.Equal(t, retBody.Message3.ContentID, "message3")
	require.Equal(t, retBody.UserData1.ContentID, "userData1")
	require.Equal(t, retBody.UserData2.ContentID, "userData2")
	require.Equal(t, retBody.UserData3.ContentID, "userData3")
	require.Equal(t, string(retBody.Message1.Content), "Hello, World!")
	require.Equal(t, string(retBody.Message2.Content), "Hello, World!!")
	require.Equal(t, string(retBody.Message3.Content), "Hello, World!!!")
	require.Equal(t, string(retBody.UserData1.Content), "username=test1&password=1234567890")
	require.Equal(t, string(retBody.UserData2.Content), "username=test2&password=1234567890")
	require.Equal(t, string(retBody.UserData3.Content), "username=test3&password=1234567890")
}
