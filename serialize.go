package openapi

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Serialize - serialize data v to corresponding media type
func Serialize(v interface{}, mediaType string) ([]byte, error) {
	var b []byte
	var err error
	switch KindOfMediaType(mediaType) {
	case MediaKindJSON:
		b, err = json.Marshal(v)
	case MediaKindXML:
		b, err = xml.Marshal(v)
	case MediaKindMultipartRelated:
		b, _, err = MultipartSerialize(v)
	default:
		return nil, errors.New("openapi client not supported serialize media type: " + mediaType)
	}
	return b, err
}
