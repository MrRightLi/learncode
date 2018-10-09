// 使用Channel来等待 goroutine的结束已经WaitGroup的使用
package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	done func()
}

func doWork(id int, worker worker) {
	for {
		for n := range worker.in {
			fmt.Printf("Work %d received %c\n",
			id, n)
			worker.done()
		}
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	worker := worker{
		in: make(chan int),
		// 这样讲wg抽象出来, 所有的事情通过这个 func 来做
		done: func() {
			wg.Done()
		},
	}
	go func() {
		doWork(id, worker)
	}()
	return worker
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker // channel of type send-only type
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	//wg.Add(20)
	for i, worker := range workers {
		wg.Add(1)
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		wg.Add(1)
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo()
}
