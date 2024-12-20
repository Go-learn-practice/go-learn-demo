package _chan

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
1. syscall.SIGINT
通常由用户通过按下 Ctrl+C 产生
**作用:**
- 通知程序终止运行。
- 典型用途是在开发者想中断程序执行时使用
**应用场景**
- 程序需要优雅地处理用户中断，例如执行清理操作、保存状态、关闭文件等

2. syscall.SIGTERM
通常由操作系统或其他进程发送，例如通过 kill 命令 (kill <PID>) 发送
**作用:**
- 请求程序终止运行，但允许程序执行清理工作
- 它是一个可以被捕获和处理的信号，程序可以选择忽略或优雅地处理
**应用场景**
- 在服务或守护进程中，收到 SIGTERM 时清理资源并安全关闭。
- 用于容器化应用（如 Docker）中的优雅关闭

3. syscall.SIGABRT
通常由程序自身发出，例如调用 abort() 函数
**作用:**
- 表示程序遇到了致命错误，需要立即中止
- 通常用于程序运行过程中检测到未恢复的异常或严重错误
**应用场景**
- 调试程序时捕获 SIGABRT，以分析导致程序中止的原因。
- 一些情况下，也用于生成核心转储（core dump）
*/

//go:generate go run ../main.go
func RunCase4Chan() {
	wg := sync.WaitGroup{}
	// 捕获信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("goroutine enter")
		<-signalChan
		os.Exit(1) // 退出当前进程
	}()

	wg.Wait()
	fmt.Println("执行结束")
}
