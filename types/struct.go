package types

import (
	"reflect"
	"strings"
)

// Map returns map from value and tag.
// the tag is optional param, not variable param.
// the value must is struct or struct pointer reference.
func Map(value interface{}, tag ...string) map[string]interface{} {
	if value == nil {
		return nil
	}

	if m, ok := value.(map[string]interface{}); ok {
		return m
	}

	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		panic("value must is struct type")
	}

	var m = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if len(tag) == 0 {
			m[field.Name] = v.Field(i).Interface()
		} else {
			tag := strings.Split(field.Tag.Get(tag[0]), ",")[0]
			if tag != "" && tag != "-" {
				m[tag] = v.Field(i).Interface()
			}
		}
	}

	return m
}
