package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generarot() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}


func worker(id int, c chan int) {
	for {
		for n := range c {
			time.Sleep(time.Second)
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

func main() {
	//var c1, c2 chan int // c1, c2 = nil

	var c1, c2 = generarot(), generarot()
	var worker = createWork(0)

	n := 0
	hasValue := false
	for {
		var activeWorker chan<- int
		if hasValue {
			activeWorker = worker
		}
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activeWorker <- n:
			hasValue = false
			//default:
			//fmt.Println("receive none")
		}
	}
}

/*
以下描述了 select 语句的语法：
	· 每个case都必须是一个通信
	· 所有channel表达式都会被求值
	· 所有被发送的表达式都会被求值
	· 如果任意某个通信可以进行，它就执行；其他被忽略。
	· 如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
	· 否则：
	· 如果有default子句，则执行该语句。
	· 如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
 */
