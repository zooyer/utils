package ioutil

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestNewReader(t *testing.T) {
	var data = []byte(`hello,world`)
	reader := NewLoopReader(data)
	for i := 0; i < 10; i++ {
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(i, string(data))
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func TestNewReaderEOF(t *testing.T) {
	var err error
	var data = []byte(`[{"key": "val"},{"key": "val"}]`)
	reader := NewLoopReaderEOF(data, false)
	decoder := json.NewDecoder(reader)
	for i := 0; i < 10; i++ {
		var obj interface{}
		if err = decoder.Decode(&obj); err != nil {
			t.Fatal(err)
		}
		t.Log(obj)
	}
}
