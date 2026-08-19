package multipart

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func testedMultipartRelatedBytes(t *testing.T, boundary string) []byte {
	jsonData := ContentDescriptor{
		MessageContainer:  &RefToBinaryData{ContentID: "message"},
		UserDataContainer: &RefToBinaryData{ContentID: "userData"},
	}

	jsonDataBytes, err := json.Marshal(jsonData)
	require.NoError(t, err)

	var reqBodyBuilder bytes.Buffer

	fmt.Fprintf(&reqBodyBuilder, "--%s\r\n", boundary)
	reqBodyBuilder.WriteString("Content-Type: application/json\r\n\r\n")
	reqBodyBuilder.Write(jsonDataBytes)
	fmt.Fprintf(&reqBodyBuilder, "\r\n--%s\r\n", boundary)
	reqBodyBuilder.WriteString("Content-Id: message\r\n")
	reqBodyBuilder.WriteString("Content-Type: application/text\r\n\r\n")
	reqBodyBuilder.WriteString("Hello, World!")
	fmt.Fprintf(&reqBodyBuilder, "\r\n--%s\r\n", boundary)
	reqBodyBuilder.WriteString("Content-Id: userData\r\n")
	reqBodyBuilder.WriteString("Content-Type: application/x-www-form-urlencoded\r\n\r\n")
	reqBodyBuilder.WriteString("username=test&password=1234567890")
	fmt.Fprintf(&reqBodyBuilder, "\r\n--%s--\r\n", boundary)

	return reqBodyBuilder.Bytes()
}
