# Golang 并发

## 1. 使用 sync.Mutex 和 sync.RWMutex
> **Mutex（互斥锁）可以保证同一时间只有一个 Goroutine 可以访问共享资源**

**示例代码：**
```go
package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock() // 加锁
	counter++
	mu.Unlock() // 解锁
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

```

## 2. 使用 sync.WaitGroup 管理 Goroutine
> **WaitGroup 用于等待一组 Goroutine 执行完毕，虽然它本身不是直接用于并发安全，但可以有效管理 Goroutine 的执行顺序**

**示例代码：**
```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 确保计数器减1
	fmt.Printf("Worker %d starting\n", id)
	// 模拟工作
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait() // 等待所有 Goroutine 完成
}

```

## 3. 使用 sync/atomic 实现原子操作
> atomic 包提供了原子操作，可以在不使用锁的情况下实现并发安全，适用于简单的计数器或标志位操作

**示例代码：**
```go
package main

import (
	"fmt"
	"sync/atomic"
)

var counter int32

func increment() {
	atomic.AddInt32(&counter, 1)
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}
	// 保证主线程等待
	fmt.Println("Final Counter:", atomic.LoadInt32(&counter))
}
```

## 4. 使用 Channel 实现并发安全
> 通过 Channel 来管理数据的访问，可以避免锁竞争，并简化代码

**示例代码：**
```go
package main

import "fmt"

func worker(jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// 启动多个 worker
	for w := 1; w <= 3; w++ {
		go worker(jobs, results)
	}

	// 发送任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}
}

```

## 使用 sync.Once 确保单次操作
> sync.Once 可以确保某个操作只执行一次，常用于初始化操作

**示例代码：**
```go
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Initializing...")
}

func main() {
	for i := 0; i < 5; i++ {
		once.Do(initialize) // 确保 initialize 只执行一次
	}
}

```

## 并发的使用场景
> 并发主要适用于以下场景：

1. I/O 密集型任务
   1. 场景：处理网络请求、文件读写、数据库操作等。
   2. 示例：高并发 HTTP 服务器、爬虫程序。
2. 计算密集型任务
   1. 场景：矩阵计算、图像处理、机器学习模型训练等。
   2. 示例：并行处理大规模数据集。
3. 生产者-消费者模型
   1. 场景：任务队列中生产任务和消费任务之间的协调。
   2. 示例：日志系统、消息队列处理。
4. 实时事件处理
   1. 场景：游戏服务器、物联网设备通信。
   2. 示例：同时处理多个客户端连接。
5. 定时或调度任务
   1. 场景：定时任务执行。
   2. 示例：实时更新服务数据的爬虫、调度系统
   
<hr>

## 总结
**根据场景选择适当的工具和模式，既能提高代码效率，也能确保线程安全**