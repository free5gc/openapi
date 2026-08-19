package encode

import (
	"encoding/json"
	"errors"
	"reflect"
)

func OneOf(v any) ([]byte, error) {
	var err error
	rv := reflect.ValueOf(v)
	if rv.IsNil() || rv.Kind() != reflect.Pointer || rv.Elem().Kind() != reflect.Struct {
		return nil, errors.New("v must be a pointer to a struct")
	}

	rv = rv.Elem()
	var data []byte

	for i := range rv.NumField() {
		field := rv.Field(i)
		if field.IsNil() {
			continue
		}

		data, err = json.Marshal(field.Interface())
		if err != nil {
			continue
		}
		return data, nil
	}

	return nil, errors.New("no value can be encoded")
}
