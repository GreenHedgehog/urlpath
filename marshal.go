package urlpath

import (
	"fmt"
	"reflect"
	"strings"
)

type Marshaler interface {
	MarshalURLValue() (value string, err error)
}

func Marshal(v interface{}) (string, error) {
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return encode("", value)
}

func MarshalScheme(scheme string, v interface{}) (string, error) {
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return encode(scheme, value)
}

func encode(scheme string, v reflect.Value) (string, error) {
	var sb strings.Builder

	fields := parseFields(scheme, v)
	for _, field := range fields {
		value, err := encodeField(field.Value)
		if err != nil {
			return "", fmt.Errorf("encode field %s failed: %v", field.tags.name, err)
		}

		if value == "" {
			value = field.tags.defaultValue
			if value == "" {
				continue
			}
		}

		sb.WriteString("/" + field.tags.name + "/" + value)
	}
	return sb.String(), nil
}

func encodeField(v reflect.Value) (string, error) {
	marshaler, ok := v.Interface().(Marshaler)
	if ok {
		return marshaler.MarshalURLValue()
	}
	return fmt.Sprintf("%v", v), nil
}
