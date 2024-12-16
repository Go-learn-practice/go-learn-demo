package one_one

import (
	"fmt"
	"producer-consumer/out"
	"sync"
)

type Task struct {
	ID int64
}

func (t *Task) run() {
	//向Out的data通道中写数据
	out.Println(t.ID)
}

var taskCh = make(chan Task, 10)

const taskNum int64 = 100

func producer(wo chan<- Task) {
	var i int64
	for i = 1; i <= taskNum; i++ {
		t := Task{
			ID: i,
		}
		wo <- t
	}
	close(wo)
}

func consumer(ro <-chan Task) {
	defer out.Close()
	for t := range ro {
		if t.ID != 0 {
			t.run()
		}
	}
}

func Exec() {
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		producer(taskCh)
	}(wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		consumer(taskCh)
	}(wg)

	wg.Wait()
	fmt.Println("执行成功")
}
