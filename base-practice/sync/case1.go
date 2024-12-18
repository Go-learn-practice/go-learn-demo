package _sync

import "fmt"

func worker(jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job * 2
	}
}

func RunChan() {
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
