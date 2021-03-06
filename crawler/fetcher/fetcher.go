package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimeiter = time.Tick(100 * time.Millisecond)
func Fetch(url string) ([]byte, error) {
	<-rateLimeiter
	resp, err := http.Get(url)
	if err != nil {
		return  nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return  nil,
			fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyRead := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyRead)
	utf8Reader := transform.NewReader(bodyRead, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)

	return all, err
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}