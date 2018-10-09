package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int)  {
	for {
		//n := <-c
		fmt.Printf("Work %d received %d\n",
			id, <-c)
	}
}

func chanDemo() {
	//var c chan int // c == nil
	c := make(chan int)

	go worker(0, c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
