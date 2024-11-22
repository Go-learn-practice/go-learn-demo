package _chan

import (
	"fmt"
	"sync"
)

/*
无缓冲通道：发送操作会阻塞，直到有 goroutine 接收数据
有缓冲通道：可以指定缓冲大小，发送操作只有在缓冲区满时才会阻塞
*/

// 创建一个无缓冲的整型通道
//var ch = make(chan int)

// 创建一个缓冲大小为 5 的通道
//var ch = make(chan int, 5)

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Produced:", i)
	}
	close(ch) // 生产完成后关闭通道
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range ch { // 消费者不断从通道中接收数据
		fmt.Println("Consumed:", val)
	}
}

func Run() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)
	wg.Wait()
}
