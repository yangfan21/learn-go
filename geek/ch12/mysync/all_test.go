package mysync

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runAllTask(id int) string {
	time.Sleep(300 * time.Millisecond)
	return fmt.Sprintf("from %d running task", id)
}

func AllResponse() string {
	numRunner := 10
	ch := make(chan string, numRunner)
	for index := 0; index < numRunner; index++ {
		go func(i int) {
			ret := runAllTask(i)
			ch <- ret
		}(index)
	}
	finalRet := ""
	for j := 0; j < numRunner; j++ {
		finalRet += <-ch + "\t"
	}
	return finalRet + "\n"
}

func TestAllResponse(t *testing.T) {
	t.Log("Before: ", runtime.NumGoroutine())
	t.Log(AllResponse())
	t.Log("After: ", runtime.NumGoroutine())
}
