package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	checkErr(err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code: ", resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	checkErr(err)

	fmt.Printf("%s\n", all)
}

func determineEncoding(r io.ReadCloser) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	checkErr(err)
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func checkErr(err error) {
	if err != err {
		panic(err)
	}
}
