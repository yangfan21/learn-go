package mysync

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("the result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for index := 0; index < numOfRunner; index++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(index)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("before: ", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(300 * time.Millisecond) // 确保所有的goroutine都执行完
	t.Log("After: ", runtime.NumGoroutine())

}
