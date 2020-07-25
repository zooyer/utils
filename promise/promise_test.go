package promise

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			p := New(func(resolve, reject func(v ...interface{})) {
				if rand.Int()%2 == 0 {
					reject("失败")
				} else {
					resolve("成功")
				}
			})

			const num = 100
			var wg sync.WaitGroup
			var ch = make(chan bool, num)
			wg.Add(num)
			for j := 0; j < num; j++ {
				p.Then(func(v ...interface{}) {
					defer wg.Done()
					ch <- true
					//t.Log(v)
				}, func(v ...interface{}) {
					defer wg.Done()
				})
			}
			wg.Wait()
			switch len(ch) {
			case 0:
			case num:
			default:
				t.Fatal("has success ye has fail")
			}
			close(ch)
		})
	}
}
