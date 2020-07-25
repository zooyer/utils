package types

import "testing"

func TestMapping(t *testing.T) {
	var s struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	tags, fields := Mapping(&s, "json")
	t.Log(tags)
	t.Log(fields)
}
