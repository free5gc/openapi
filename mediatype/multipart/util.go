package multipart

import (
	"reflect"
)

func IfMatchIdSetContent(contentId string, target *[]byte, part *RelatedContent) {
	if part == nil || target == nil {
		return
	}
	if part.ContentID == contentId {
		*target = part.Content
	}
}

func GetRelatedContentFields(v any, contentType string) []reflect.Value {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	valType := val.Type()
	var result []reflect.Value

	for i := range valType.NumField() {
		field := valType.Field(i)
		if field.Type == relatedContentType {
			ctTag := parseMultipartFieldParameters(field.Tag.Get("multipart"))
			if ctTag == contentType {
				result = append(result, val.Field(i))
			}
		}
	}
	return result
}

func MatchRelatedContent(contentId string, fields []reflect.Value) []byte {
	for _, field := range fields {
		if field.Type() == relatedContentType {
			rc := field.Interface().(*RelatedContent)
			if rc != nil && rc.ContentID == contentId {
				return rc.Content
			}
		}
	}
	return nil
}
