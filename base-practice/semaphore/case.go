package _semaphore

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"math/rand"
	"sync"
	"time"
)

/*
1. 创建信号量：semaphore.NewWeighted(3) 创建了一个容量为 3 的信号量，这意味着最多只能有 3 个 goroutine 同时运行
2. Acquire()：在每个 goroutine 中，通过 sem.Acquire(context.Background(), 1) 请求一个信号量。这里的 1 表示每个 goroutine 请求 1 个信号量。如果没有可用的信号量，goroutine 会阻塞，直到有信号量可用
3. Release()：sem.Release(1) 释放信号量，允许其他 goroutine 获取信号量继续执行
4. 工作负载：每个 goroutine 模拟一些工作（time.Sleep），并在工作完成后释放信号量
*/

func RunSem() {
	// 创建一个信号量，最大容量为 3
	sem := semaphore.NewWeighted(3)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// 请求信号量
			if err := sem.Acquire(context.Background(), 1); err != nil {
				fmt.Printf("Goroutine %d 无法获取信号量: %v\n", i, err)
				return
			}

			defer sem.Release(1)

			// 模拟处理任务
			fmt.Printf("Goroutine %d 开始工作\n", i)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Printf("Goroutine %d 完成工作\n", i)
		}(i)
	}

	// 等待所有 goroutine 完成
	wg.Wait()
}
