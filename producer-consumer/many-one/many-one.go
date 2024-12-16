package many_one

import (
	"fmt"
	"producer-consumer/out"
	"sync"
)

type Task struct {
	ID int64
}

func (t *Task) run() {
	out.Println(t.ID)
}

var taskCh = make(chan Task, 10)

const taskNum int64 = 100
const nums int64 = 100

func producer(wo chan<- Task, startNum int64, nums int64) {
	var i int64
	for i = startNum; i <= startNum+nums; i++ {
		t := Task{
			ID: i,
		}
		//fmt.Println("生产者写入数据")
		wo <- t
	}
}

func consumer(ro <-chan Task) {
	defer func() {
		out.Close()
	}()
	for t := range ro {
		if t.ID != 0 {
			t.run()
		}
	}
}

func Exec() {
	wg := &sync.WaitGroup{}
	cwg := &sync.WaitGroup{}
	var i int64
	for i = 0; i < taskNum; i += nums {
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()
			producer(taskCh, i, nums)
		}(i)
	}

	go func() {
		wg.Wait()
		close(taskCh)
	}()

	cwg.Add(1)
	go func() {
		defer cwg.Done()
		consumer(taskCh)
	}()
	cwg.Wait()
	fmt.Println("执行成功")
}
