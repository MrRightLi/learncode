package main

import (
	"fmt"
	"time"
)

type atomicInt int

func (a *atomicInt) increament() {
	*a++
}

func (a *atomicInt) get() int {
	return int(*a)
}

func main() {
	var a atomicInt
	a.increament()
	go func() {
		a.increament()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
