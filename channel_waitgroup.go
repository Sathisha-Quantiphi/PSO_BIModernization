package main

import (
	"fmt"
	"sync"
	"time"
)

func read() (elemChan chan int) {
	elemChan = make(chan int)
	go func() {
		for k := 0; k < 1000; k++ {
			elemChan <- k
		}
		close(elemChan)
	}()
	return
}

func main() {
	elemChan := read()
	wg := sync.WaitGroup{}
	for k := 0; k < 2; k++ {
		wg.Add(1)
		go func(k int) {
			for {
				e, more := <-elemChan
				if !more {
					wg.Done()
					return
				}
				fmt.Printf("goroutine #%d: %d\n", k, e)
				time.Sleep(1000 * time.Nanosecond)
			}
		}(k)
	}
	wg.Wait()
}
