package multipart

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRelatedContentFields(t *testing.T) {
	m := MultipartRelatedTestBody{
		JsonData: &ContentDescriptor{
			MessageContainer: &RefToBinaryData{
				ContentID: "message",
			},
			UserDataContainer: &RefToBinaryData{
				ContentID: "userData",
			},
		},
		Message: &RelatedContent{
			ContentID: "message",
			Content:   []byte("message"),
		},
		UserData: &RelatedContent{
			ContentID: "userData",
			Content:   []byte("userData"),
		},
	}
	textFields := GetRelatedContentFields(m, "application/text")
	formFields := GetRelatedContentFields(m, "application/x-www-form-urlencoded")
	require.Equal(t, len(textFields), 1)
	require.Equal(t, len(formFields), 1)
	require.Equal(t, textFields[0].Interface().(*RelatedContent).ContentID, "message")
	require.Equal(t, formFields[0].Interface().(*RelatedContent).ContentID, "userData")
}

func TestMatchRelatedContent(t *testing.T) {
	m := MultipartRelatedTestBody{
		JsonData: &ContentDescriptor{
			MessageContainer: &RefToBinaryData{
				ContentID: "message",
			},
			UserDataContainer: &RefToBinaryData{
				ContentID: "userData",
			},
		},
		Message: &RelatedContent{
			ContentID: "message",
			Content:   []byte("message"),
		},
		UserData: &RelatedContent{
			ContentID: "userData",
			Content:   []byte("userData"),
		},
	}
	textFields := GetRelatedContentFields(m, "application/text")
	formFields := GetRelatedContentFields(m, "application/x-www-form-urlencoded")
	content := MatchRelatedContent("message", textFields)
	require.Equal(t, content, []byte("message"))
	content = MatchRelatedContent("userData", formFields)
	require.Equal(t, content, []byte("userData"))
}
