package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/list/"):] // /list/abc.txt
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()


	})
}
