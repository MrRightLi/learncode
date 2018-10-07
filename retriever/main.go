package main

import (
	"fmt"
	"project/google_code/retriever/mock"
	"project/google_code/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is google.com"}
	fmt.Println(download(r))

	r = real.Retriever{}
	fmt.Println(r.Get("http://www.yufu365.com"))
}
