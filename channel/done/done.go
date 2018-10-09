package main

import (
	"fmt"
)

type worker struct {
	in chan int
	done chan bool
}

func doWork(id int, worker worker) {
	for {
		for n := range worker.in {
			fmt.Printf("Work %d received %c\n",
			id, n)
			worker.done <- true
		}
	}
}

func createWorker(id int) worker {
	worker := worker{
		in: make(chan int),
		done: make(chan bool),
	}
	go func() {
		doWork(id, worker)
	}()

	return worker
}

func chanDemo() {
	var workers [10]worker // channel of type send-only type
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		fmt.Printf("发送: %c\n", 'a' + i)
		//<-workers[i].done
		//fmt.Printf("收到%c的done\n", 'a' + i)

	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		fmt.Printf("发送: %c\n", 'A' + i)
		//<-workers[i].done
		//fmt.Printf("收到%c的done\n", 'A' + i)
	}

	// wait for all of them
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func main() {
	chanDemo()
}
