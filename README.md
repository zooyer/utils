# utils - golang library

![](https://travis-ci.org/boennemann/badges.svg?branch=master)  ![](https://img.shields.io/badge/license-MIT-blue.svg)  ![](https://img.shields.io/badge/godoc-reference-blue.svg)

Golang common uitls tools.

![](https://github.com/golang/go/blob/master/doc/gopher/fiveyears.jpg?raw=true)



#### Download and Install

```shell
go get github.com/zooyer/utils
```



#### Features

- loop reader.
- deep copy and compared.
- bytes to string and string to bytes is zero memory.
- struct to map from tags.
- struct to tag and field mapping.



#### Example

1. loop reader

   ```go
   package main
   
   import (
   	"fmt"
   	"io/ioutil"
   
   	"github.com/zooyer/utils/ioutils"
   )
   
   func main() {
   	var data = []byte(`hello,world`)
   	reader := ioutils.NewLoopReader(data)
   	for i := 0; i < 10; i++ {
   		data, err := ioutil.ReadAll(reader)
   		if err != nil {
   			panic(err)
   		}
   		fmt.Println(i, string(data))
   	}
   	data, err := ioutil.ReadAll(reader)
   	if err != nil {
   		panic(err)
   	}
   	fmt.Println(string(data))
   }
   ```

2. deep copy and compared

   ```go
   package main
   
   import (
   	"fmt"
   	
   	"github.com/zooyer/utils/types"
   )
   
   func main() {
   	type Type struct {
   		Map     map[string]*int
   		Int     int
   		Pointer *int
   	}
   
   	var i = 100
   	var obj1 = Type{
   		Map: map[string]*int{
   			"int": &i,
   		},
   		Int:     123,
   		Pointer: nil,
   	}
   
   	var obj2, ok = types.Clone(obj1).(Type)
   	if !ok {
   		panic(fmt.Sprintf("%#v", types.Clone(obj1)))
   	}
   
   	var change = func(typ *Type) {
   		var i = 999
   		typ.Pointer = &i
   		typ.Int = i
   		typ.Map["int"] = &i
   	}
   
   	change(&obj1)
   	if types.Equal(obj1, obj2) {
   		panic("clone failed")
   	}
   
   	change(&obj2)
   	if !types.Equal(obj1, obj2) {
   		panic("clone failed")
   	}
   }
   ```

3. bytes to string and string to bytes

   ```go
   package main
   
   import (
   	"fmt"
   
   	"github.com/zooyer/utils/types"
   )
   
   func main() {
   	var str = "abc"
   	b := types.Bytes(str)
   	fmt.Println("bytes(read only):", b)
   	b = []byte("abc")
   
   	str = types.String(b)
   	b[1] = 'a'
   	b[2] = 'a'
   
   	if str != "aaa" {
   		panic(fmt.Sprintf("str != aaa: %v", str))
   	}
   }
   ```

4. struct to map

   ```go
   package main
   
   import (
   	"github.com/zooyer/utils/types"
   )
   
   func main() {
   	var people = struct {
   		Name string `key:"name"`
   		Age  int    `key:"age"`
   	}{
   		Name: "test",
   		Age:  2018,
   	}
   
   	maps := types.Map(people)
   	if types.Equal(maps, map[string]interface{}{"Name": "string", "Age": 2018}) {
   		panic("TestMap failed")
   	}
   
   	maps = types.Map(people, "key")
   	if types.Equal(maps, map[string]interface{}{"name": "string", "age": 2018}) {
   		panic("TestMap failed")
   	}
   }
   ```

5. struct to tag and field mapping

   ```go
   package main
   
   import (
   	"fmt"
   	"github.com/zooyer/utils/types"
   )
   
   func main() {
   	var s struct {
   		Name string `json:"name"`
   		Age  int    `json:"age"`
   	}
   	tags, fields := types.Mapping(&s, "json")
   	fmt.Println(tags)
   	fmt.Println(fields)
   }
   ```

   