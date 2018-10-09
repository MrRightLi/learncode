package main

import (
	"fmt"
	"project/google_code/queue"

	"golang.org/x/tools/container/intsets"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	if q.IsEmpty() == false {
		fmt.Println(q.Pop())
	}

	fmt.Println(q.IsEmpty())

	testSparse()
}

func testSparse() {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(100000))
}
