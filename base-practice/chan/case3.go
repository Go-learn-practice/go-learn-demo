package _chan

import "fmt"

func Run3() {
	ch := make(chan int, 10)
	done := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				ch <- i*10 + j
			}
		}
		close(done) // 通知消费者任务完成
	}()

	go func() {
		<-done    // 等待生产完成
		close(ch) // 消费完成后关闭通道
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
