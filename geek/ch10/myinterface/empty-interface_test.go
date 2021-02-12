package myinterface

import (
	"fmt"
	"testing"
)

// 倾向于使用更小的接口， 很多接口只有一个方法

//较大的接口 可以由多个小的接口组合而成

// 只依赖必要功能的最小接口

type myInterface interface {
}

func DoSomething(i interface{}) {
	//if p, ok := i.(int); ok {
	//	fmt.Printf(" Integer and value: %d\n", p)
	//	return
	//}
	//if p, ok := i.(string); ok {
	//	fmt.Printf(" String and value: %s\n", p)
	//	return
	//}
	switch p := i.(type) {
	case int:
		fmt.Printf(" Integer and value: %d\n", p)
	case string:
		fmt.Printf(" String and value: %s\n", p)
	case bool:
		fmt.Printf(" Bool and value: %v\n", p)
	default:
		fmt.Printf("Unknow and value: %T\n", i)
	}
}

func TestEmpty(t *testing.T) {
	var a myInterface = 12
	DoSomething(a)

}
