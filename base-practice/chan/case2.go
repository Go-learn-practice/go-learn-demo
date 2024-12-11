package _chan

import (
	"fmt"
	"sync"
)

func Run2() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println(id)
		}(i)
	}
	wg.Wait()
}
