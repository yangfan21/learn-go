package mysync

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type SingletonObject struct {
	Data string
}

var singleInstance *SingletonObject
var once sync.Once

func GetSingleObj() *SingletonObject {
	once.Do(func() {
		fmt.Println("create single object")
		singleInstance = new(SingletonObject)
	})
	return singleInstance
}

func TestGetSingleObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingleObj()
			fmt.Printf("obj address: %x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
