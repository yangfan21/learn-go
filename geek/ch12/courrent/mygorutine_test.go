package courrent

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	// 不嫩如下的原因是index 虽然在goroutine中是可用的，但他是被多个goroutine 共享的变量，就需要用锁的机制实现
	//for index := 1; index < 10; index++ {
	//	go func() {
	//		fmt.Println(index)
	//	}()
	//}

	// go的方法函数参数传递是值传递，都是拷贝了一份index 给到i, 每个index是独立的，不存在竞争关系
	for index := 1; index < 10; index++ {
		go func(i int) {
			fmt.Println(i)
		}(index)
	}
	time.Sleep(50 * time.Millisecond)
}

// counter 在不同的goroutine中做自增，会产生竞争，这就不是线程安全的，需要用锁对共享的内存进行处理
func TestMyLock(t *testing.T) {
	counter := 0
	for index := 0; index < 5000; index++ {
		go func(i int) {
			counter++
			fmt.Println(i)
		}(index)
	}
	time.Sleep(1 * time.Second)
	t.Logf("get counter value: %d", counter)
}

func TestSafeLock(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for index := 0; index < 5000; index++ {
		go func(i int) {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			fmt.Println(i)
		}(index)
	}
	time.Sleep(2 * time.Second)
	t.Logf("get counter value: %d", counter)
}

// 这样操作并没有改变counter被竞争的事实
func TestFailGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	counter := 0
	for index := 0; index < 5000; index++ {
		go func() {
			defer func() {
				wg.Done()
			}()
			wg.Add(1)
			counter++
		}()
	}
	wg.Wait()
	t.Logf("get counter value: %d", counter)
}

func TestSafeGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	var mut sync.Mutex
	counter := 0
	for index := 0; index < 5000; index++ {
		go func() {
			defer func() { wg.Done() }()
			defer func() { mut.Unlock() }()
			wg.Add(1)
			mut.Lock()
			counter++
		}()
	}
	wg.Wait()
	t.Logf("get counter value: %d", counter)
}
