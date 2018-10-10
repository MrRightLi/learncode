package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		for n := range c {
			fmt.Printf("Work %d received %d\n",
			id, n)
		}
	}
}

func createWork(id int) chan<- int {
	channel := make(chan int)
	go func() {
		worker(id, channel)
	}()

	return channel
}

func chanDemo() {
	var channels [10]chan<- int // channel of type send-only type
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

func bufferedChannel() {
	c := make(chan int, 3) // buffer channel 比较省资源
	go worker(0, c)
	c <- 'a' + 0
	c <- 'a' + 1
	c <- 'a' + 2
	c <- 'a' + 3
	c <- 'a' + 4
	time.Sleep(time.Millisecond)
}

func closeChannel() {
	c := make(chan int)
	//c := make(chan int, 3)
	go worker(0, c)
	c <- 'a' + 0
	c <- 'a' + 1
	c <- 'a' + 2
	c <- 'a' + 3
	c <- 'a' + 4
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Printf("Channel as first-class citizen")
	chanDemo()

	fmt.Println("Buffered channel")
	bufferedChannel()

	fmt.Println("Channel close and range")
	closeChannel()

	// 当然在channel的发送者不发送之xiaojie后,接受也就随着关闭
}
