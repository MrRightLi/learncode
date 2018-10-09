package main

import (
	"fmt"
	"time"
)

func main() {
	var a [1000]int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			//a[ii]++
			//runtime.Gosched()
			fmt.Println("Hello from goroutine %d \n", i)
		}(i)
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
}
