package _sync

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var once sync.Once

func initialize() {
	fmt.Println("Initializing...")
}

func RunOnce() {
	for i := 0; i < 5; i++ {
		once.Do(initialize) // 确保 initialize 只执行一次
	}
}

var counter int32

func increment() {
	atomic.AddInt32(&counter, 1)
}

func RunAtomic() {
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
