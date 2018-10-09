package main

import (
	"fmt"
	"time"
)

func createWork(id int) chan<- int {


	channel := make(chan int)
	go func() {
		for {
			fmt.Printf("Work %d received %c\n",
				id, <-channel)
		}
	}()

	return channel
}

func chanDemo() {
	var channels [10]chan<- int  // channel of type send-only type
	for i := 0; i < 10; i++ {
		channels[i] = createWork(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
