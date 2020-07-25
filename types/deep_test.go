package types

import (
	"fmt"
	"testing"
)

func TestClone(t *testing.T) {
	type Type struct {
		Map     map[string]*int
		Int     int
		Pointer *int
	}

	var i = 100
	var obj1 = Type{
		Map: map[string]*int{
			"int": &i,
			//"int": nil, // gob无法对这个指针做映射
		},
		Int:     123,
		Pointer: nil,
	}

	var obj2, ok = Clone(obj1).(Type)
	if !ok {
		t.Fatal(fmt.Sprintf("%#v", Clone(obj1)))
	}

	var change = func(typ *Type) {
		var i = 999
		typ.Pointer = &i
		typ.Int = i
		typ.Map["int"] = &i
	}

	change(&obj1)
	if Equal(obj1, obj2) {
		t.Fatal("clone failed")
	}

	change(&obj2)
	if !Equal(obj1, obj2) {
		t.Fatal("clone failed")
	}
}
