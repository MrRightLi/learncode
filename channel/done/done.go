package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func doWork(id int, worker worker) {
	for {
		for n := range worker.in {
			fmt.Printf("Work %d received %c\n",
			id, n)
			worker.wg.Done()
		}
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	worker := worker{
		in: make(chan int),
		wg: wg,
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
