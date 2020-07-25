package types

import (
	"bytes"
	"encoding/gob"
	"reflect"
	"sync"
)

var (
	mutex     sync.RWMutex
	registers = make(map[string]bool)
)

// register register value to gob.
func register(value interface{}) {
	if value == nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	// Default to printed representation for unnamed types
	rt := reflect.TypeOf(value)
	name := rt.String()

	// But for named types (or pointers to them), qualify with import path (but see inner comment).
	// Dereference one pointer looking for a named type.
	star := ""
	if rt.Name() == "" {
		if pt := rt; pt.Kind() == reflect.Ptr {
			star = "*"
			// NOTE: The following line should be rt = pt.Elem() to implement
			// what the comment above claims, but fixing it would break compatibility
			// with existing gobs.
			//
			// Given package p imported as "full/p" with these definitions:
			//     package p
			//     type T1 struct { ... }
			// this table shows the intended and actual strings used by gob to
			// name the types:
			//
			// Type      Correct string     Actual string
			//
			// T1        full/p.T1          full/p.T1
			// *T1       *full/p.T1         *p.T1
			//
			// The missing full path cannot be fixed without breaking existing gob decoders.
			rt = pt
		}
	}
	if rt.Name() != "" {
		if rt.PkgPath() == "" {
			name = star + rt.Name()
		} else {
			name = star + rt.PkgPath() + "." + rt.Name()
		}
	}

	if !registers[name] {
		gob.Register(value)
		registers[name] = true
	}
}

// Clone returns a new value from deep clone value.
func Clone(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	register(value)

	var err error
	var buf bytes.Buffer
	if err = gob.NewEncoder(&buf).Encode(value); err != nil {
		return nil
	}
	var cloned = reflect.New(reflect.TypeOf(value))
	if err = gob.NewDecoder(&buf).Decode(cloned.Interface()); err != nil {
		return nil
	}

	return cloned.Elem().Interface()
}

// Equal returns deep a equal b.
func Equal(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}
