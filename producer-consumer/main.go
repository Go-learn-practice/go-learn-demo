package main

import (
	"fmt"
	"os"
	"os/signal"
	many_many "producer-consumer/many-many"
	"producer-consumer/out"
	"syscall"
)

func main() {
	//wg := sync.WaitGroup{}

	o := out.NewOut()

	//wg.Add(1)
	go func() {
		//defer wg.Done()
		fmt.Println("开启协程读取数据，初始时通道中未有数据会阻塞")
		o.OutPut()
	}()

	//one_one.Exec()
	//one_many.Exec()
	//many_one.Exec()

	many_many.Exec()

	//wg.Wait()

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	<-sign
}
