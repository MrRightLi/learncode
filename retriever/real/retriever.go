package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	checkErr(err)
	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	checkErr(err)

	return string(result)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


