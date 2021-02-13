package mychannel

import (
	"fmt"
	"sync"
	"testing"
)

func DataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		defer close(ch)
		for index := 0; index < 10; index++ {
			ch <- index
		}
		wg.Done()
	}()
}

func DataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Printf("from channel data is: %d\n", data)
			} else {
				break
			}

		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	DataProducer(ch, &wg)
	wg.Add(1)
	DataReceiver(ch, &wg)
	wg.Wait()
}
