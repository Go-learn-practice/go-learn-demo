package many_many

import (
	"fmt"
	"producer-consumer/out"
	"time"
)

type Task struct {
	ID int64
}

func (t *Task) run() {
	out.Println(t.ID)
}

var taskCh = make(chan Task, 10)
var done = make(chan struct{})

const taskNum int64 = 100

// 生产者
func producer(wo chan<- Task, done <-chan struct{}) {
	var i int64
	for {
		if i >= taskNum {
			i = 0
		}
		i++
		t := Task{ID: i}
		select {
		case wo <- t:
		case <-done:
			out.Println("生产者退出")
			return
		}
	}
}

// 消费者
func consumer(ro <-chan Task, done <-chan struct{}) {
	for {
		select {
		case t, _ := <-ro:
			if t.ID != 0 {
				t.run()
			}
		case <-done:
			for t := range ro {
				if t.ID != 0 {
					t.run()
				}
			}
			out.Println("消费者退出")
			return
		}
	}
}

func Exec() {
	go producer(taskCh, done)
	go producer(taskCh, done)
	go producer(taskCh, done)
	go producer(taskCh, done)
	go producer(taskCh, done)
	go producer(taskCh, done)

	go consumer(taskCh, done)
	go consumer(taskCh, done)
	go consumer(taskCh, done)
	go consumer(taskCh, done)

	time.Sleep(time.Second * 5)
	close(done)
	close(taskCh)
	time.Sleep(time.Second * 5)
	fmt.Println(len(taskCh))
}
