package openapi

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"mime"

	"github.com/free5gc/openapi/mediatype"
	"github.com/free5gc/openapi/mediatype/multipart"
)

// Serialize - serialize data v to corresponding media type
// return contentType, bodyBytes, error
func Serialize(v any, mediaType string) (string, []byte, error) {
	var b []byte
	var err error
	var contentType string
	switch mediatype.KindOfMediaType(mediaType) {
	case mediatype.MediaKindJSON:
		b, err = json.Marshal(v)
		contentType = mediaType
	case mediatype.MediaKindXML:
		b, err = xml.Marshal(v)
		contentType = mediaType
	case mediatype.MediaKindMultipartRelated:
		var boundary string
		boundary, b, err = multipart.Marshal(v)
		if err != nil {
			return "", nil, err
		}
		contentType = mime.FormatMediaType(
			"multipart/related", map[string]string{"boundary": boundary})
	default:
		return "", nil, errors.New("openapi client not supported serialize media type: " + mediaType)
	}
	return contentType, b, err
}
