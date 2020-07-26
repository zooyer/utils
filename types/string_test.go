package types

import "testing"

func TestToBytes(t *testing.T) {
	var str = "abc"
	b := Bytes(str)
	t.Log("bytes(read only):", b)
}

func TestToString(t *testing.T) {
	var b = []byte("abc")
	str := String(b)
	b[1] = 'a'
	b[2] = 'a'

	if str != "aaa" {
		t.Fatal("str != aaa:", str)
	}
}
