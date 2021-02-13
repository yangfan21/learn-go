package mychannel

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(100 * time.Millisecond)
	return "service done !"
}

func AsyncService() chan string {
	retChan := make(chan string, 10)
	go func() {
		ret := service()
		fmt.Println("return service func result.")
		retChan <- ret
		fmt.Println("service exit.")
	}()
	return retChan
}

func otherTask() string {
	fmt.Println("working in other function ")
	time.Sleep(20 * time.Millisecond)
	return "other task done !"
}

func TestService(t *testing.T) {
	t.Log(service())
	t.Log(otherTask())
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	t.Logf("get service data: %s\n", <-retCh)
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Logf("get func data: %s", ret)
	case <-time.After(800 * time.Millisecond):
		t.Error("time out to call service.")
	}

}
