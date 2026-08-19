package decode

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
)

func OneOf(data []byte, v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() || rv.Elem().Kind() != reflect.Struct {
		return errors.New("v must be a pointer to a struct")
	}

	rv = rv.Elem()

	// for each field in the struct, and try to unmarshal the data into it
	for i := range rv.NumField() {
		field := rv.Field(i)

		dec := json.NewDecoder(bytes.NewReader(data))
		dec.DisallowUnknownFields()

		// New value of the field
		fv := reflect.New(field.Type()).Interface()

		err := dec.Decode(fv)
		if err == nil {
			field.Set(reflect.ValueOf(fv).Elem())
			return nil
		}
	}

	return errors.New("unmarshal error for oneof field")
}
