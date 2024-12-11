package _chan

import (
	"fmt"
	"sync"
)

func RunCase() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println("生产者协程")
			defer wg.Done()
			for j := 0; j < 5; j++ {
				ch <- id*10 + j // 向通道发送数据
			}
		}(i)
	}

	go func() {
		fmt.Println("enter关闭协程")
		wg.Wait() // 等待所有生产者完成
		close(ch) // 关闭通道，通知消费者
	}()

	cwg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		cwg.Add(1)
		go func() {
			fmt.Println("消费者协程")
			defer cwg.Done()
			for val := range ch { // 从通道接收数据
				fmt.Println(val)
			}
		}()
	}

	fmt.Println("主协程")
	cwg.Wait() // 等待所有生产者和消费者完成
}
