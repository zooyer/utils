package types

import (
	"reflect"
	"unsafe"
)

// Bytes returns bytes from str(the bytes read only).
func Bytes(str string) (bytes []byte) {
	b := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	s := (*reflect.StringHeader)(unsafe.Pointer(&str))
	b.Data = s.Data
	b.Len = s.Len
	b.Cap = s.Len
	return
}

// String returns str from bytes(bytes changes affect the str).
func String(bytes []byte) (str string) {
	b := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	s := (*reflect.StringHeader)(unsafe.Pointer(&str))
	s.Data = b.Data
	s.Len = b.Len
	return
}
