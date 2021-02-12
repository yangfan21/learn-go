package myfunc

import (
	"fmt"
	"testing"
	"time"
)

func TestParamFunc(t *testing.T) {
	spent := timeSpent(slowFunc)
	t.Logf("the result: %v", spent(20))
}

// 基础函数
func slowFunc(op int) int {
	time.Sleep(2 * time.Second)
	return op
}

// 可扩展基础函数的功能 入参与出餐都是一个函数，内部调用这个函数
func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("the func time spent: ", time.Since(start).Seconds())
		return ret
	}
}
