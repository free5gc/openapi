package multipart

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"reflect"
	"slices"
	"strings"

	"github.com/free5gc/openapi/mediatype"
	"github.com/pkg/errors"
)

type RelatedContent struct {
	ContentID string
	Content   []byte
}

var (
	relatedContentType = reflect.TypeOf((*RelatedContent)(nil))
)

// unmarshalStructFieldMapping is a map of struct fields to their corresponding field index in the multipart struct
type unmarshalStructFieldMapping struct {
	jsonData int
	// key is the contentType, value is the field index
	otherContents map[string][]int
}

func Unmarshal(boundary string, data []byte, v any) error {
	r := multipart.NewReader(bytes.NewReader(data), boundary)

	fieldMapping, err := resolveStructFieldMapping(v)
	if err != nil {
		return fmt.Errorf("failed to resolve struct field mapping for %T: %w", v, err)
	}
	val := reflect.Indirect(reflect.ValueOf(v))

	idx := 1
	for {
		p, err := r.NextPart()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}
		contentPayload, err := io.ReadAll(p)
		if err != nil {
			return err
		}

		contentType := p.Header.Get("Content-Type")
		contentID := p.Header.Get("Content-ID")

		if contentType == "" {
			return fmt.Errorf("content type is empty for content %d", idx)
		}

		switch mediatype.KindOfMediaType(contentType) {
		case mediatype.MediaKindJSON:
			fieldVal := val.Field(fieldMapping.jsonData)
			if fieldVal.Type().Kind() != reflect.Pointer {
				fieldVal = fieldVal.Addr()
			}

			if fieldVal.IsNil() {
				fieldVal.Set(reflect.New(fieldVal.Type().Elem()))
			}

			err = json.Unmarshal(contentPayload, fieldVal.Interface())
			if err != nil {
				return fmt.Errorf("failed to unmarshal content %d: %w", idx, err)
			}
		default:
			if contentID == "" {
				return fmt.Errorf("contentID is required for content %d of type %s", idx, contentType)
			}
			// pop the first field index
			fieldIndices, ok := fieldMapping.otherContents[contentType]
			if !ok || len(fieldIndices) == 0 {
				return fmt.Errorf("no field found for content type %s", contentType)
			}
			fieldIdx := fieldIndices[0]
			fieldMapping.otherContents[contentType] = slices.Delete(fieldIndices, 0, 1)

			fieldVal := val.Field(fieldIdx)
			if fieldVal.Type() == relatedContentType {
				fieldVal.Set(reflect.ValueOf(&RelatedContent{ContentID: contentID, Content: contentPayload}))
			}
		}

		idx++
	}
	return nil
}

func resolveStructFieldMapping(v any) (*unmarshalStructFieldMapping, error) {
	val := reflect.Indirect(reflect.ValueOf(v))
	valType := val.Type()

	fieldMapping := &unmarshalStructFieldMapping{
		jsonData:      0,
		otherContents: make(map[string][]int),
	}

	for i := range valType.NumField() {
		fieldType := valType.Field(i)

		contentType := parseMultipartFieldParameters(fieldType.Tag.Get("multipart"))

		if contentType == "" {
			return nil, fmt.Errorf("content type is empty for field %s", fieldType.Name)
		}

		kind := mediatype.KindOfMediaType(contentType)

		switch kind {
		case mediatype.MediaKindJSON:
			fieldMapping.jsonData = i
		default:
			fieldMapping.otherContents[contentType] = append(fieldMapping.otherContents[contentType], i)
		}
	}
	return fieldMapping, nil
}

func Marshal(v any) (string, []byte, error) {
	var err error

	val := reflect.Indirect(reflect.ValueOf(v))
	valType := val.Type()
	buf := bytes.NewBuffer(nil)
	w := multipart.NewWriter(buf)
	boundary, err := generateBoundary()
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate boundary: %w", err)
	}
	err = w.SetBoundary(boundary)
	if err != nil {
		return "", nil, fmt.Errorf("failed to set boundary: %w", err)
	}

	for i := range val.NumField() {
		field := val.Field(i)
		if field.IsNil() {
			continue
		}

		var mimeHeader = textproto.MIMEHeader{}
		var content []byte

		contentType := parseMultipartFieldParameters(val.Type().Field(i).Tag.Get("multipart"))

		kind := mediatype.KindOfMediaType(contentType)

		switch kind {
		case mediatype.MediaKindJSON:
			mimeHeader.Set("Content-Type", "application/json")
			content, err = json.Marshal(field.Interface())
			if err != nil {
				return "", nil, err
			}
		default:
			// check if the field is `RelatedContent`
			if field.IsNil() || field.Type() != relatedContentType {
				return "", nil, fmt.Errorf("invalid field for multipart: %s", valType.Field(i).Name)
			}

			relatedContent := field.Interface().(*RelatedContent)
			if relatedContent.ContentID == "" {
				return "", nil, errors.New("contentID is required")
			}
			mimeHeader.Set("Content-Type", contentType)
			mimeHeader.Set("Content-ID", relatedContent.ContentID)
			content = relatedContent.Content
		}

		var part io.Writer
		part, err = w.CreatePart(mimeHeader)
		if err != nil {
			return "", nil, err
		}
		_, err = part.Write(content)
		if err != nil {
			return "", nil, err
		}
	}

	err = w.Close()
	if err != nil {
		return "", nil, err
	}

	return boundary, buf.Bytes(), nil
}

func parseMultipartFieldParameters(str string) (contentType string) {
	for _, part := range strings.Split(str, ",") {
		switch {
		case strings.HasPrefix(part, "contentType:"):
			contentType = part[12:]
		}
	}
	return
}
