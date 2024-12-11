package _ctx

import (
	"context"
	"fmt"
	"time"
)

func processTask(ctx context.Context, taskName string, duration time.Duration) {
	select {
	case <-time.After(duration):
		// 模拟任务执行
		fmt.Printf("%s: Task completed after %v\n", taskName, duration)
	case <-ctx.Done():
		// 当 context 被取消时，任务应该停止
		fmt.Printf("%s: Task cancelled: %v\n", taskName, ctx.Err())
	}
}

func Run() {
	ctx, cancel := context.WithCancel(context.Background())

	// 启动多个 goroutine 来执行任务
	go processTask(ctx, "Task 1", 3*time.Second) // 3 秒的任务
	go processTask(ctx, "Task 2", 5*time.Second) // 5 秒的任务
	go processTask(ctx, "Task 3", 7*time.Second) // 7 秒的任务

	// 模拟 4 秒后取消所有任务
	time.Sleep(4 * time.Second)
	cancel() // 调用 cancel 函数取消所有任务

	// 等待所有任务完成
	time.Sleep(1 * time.Second) // 让所有 goroutine 有时间输出结果
}
