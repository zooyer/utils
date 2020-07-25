package strutil

import (
	"reflect"
	"unsafe"
)

// Bytes 字符串转字节切片(注:此切片只读)
func Bytes(str string) []byte {
	var bytes struct {
		addr uintptr
		len  int
		cap  int
	}

	var header = *(*reflect.StringHeader)(unsafe.Pointer(&str))
	bytes.len = len(str)
	bytes.cap = len(str)
	bytes.addr = header.Data

	return *(*[]byte)(unsafe.Pointer(&bytes))
}

// String 字节切片转字符串(注:修改切片内容会影响字符串)
func String(b []byte) string {
	var str reflect.StringHeader
	if len(b) > 0 {
		str.Data = uintptr(unsafe.Pointer(&b[0]))
	}
	str.Len = len(b)

	return *(*string)(unsafe.Pointer(&str))
}
