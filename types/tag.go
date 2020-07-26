package types

import (
	"reflect"
	"strings"
)

// Mapping returns tags map and fields map from value and tag.
func Mapping(value interface{}, tag string) (tags, fields map[string]string) {
	if value == nil {
		return nil, nil
	}

	var fieldName, tagName string

	t := reflect.TypeOf(value)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("value must is struct type")
	}

	tags = make(map[string]string)
	fields = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		tagName = t.Field(i).Tag.Get(tag)
		tagName = strings.Split(tagName, ",")[0]
		fieldName = t.Field(i).Name
		if len(tagName) != 0 && tagName != "-" {
			tags[tagName] = fieldName
			fields[fieldName] = tagName
		}
	}

	return tags, fields
}
