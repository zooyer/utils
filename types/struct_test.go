package types

import "testing"

func TestMap(t *testing.T) {
	var people = struct {
		Name string `key:"name"`
		Age  int    `key:"age"`
	}{
		Name: "test",
		Age:  2018,
	}

	maps := Map(people)
	if Equal(maps, map[string]interface{}{"Name": "string", "Age": 2018}) {
		t.Fatal("TestMap failed")
	}

	maps = Map(people, "key")
	if Equal(maps, map[string]interface{}{"name": "string", "age": 2018}) {
		t.Fatal("TestMap failed")
	}
}
