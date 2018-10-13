package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contens, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contens)

	const resultSize = 470

	fmt.Println(len(result.Request))

	if len(result.Request) != resultSize {
		t.Errorf("result should have %d" +
			"requests;but had %d", resultSize, len(result.Request))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d" +
			"requests;but had %d", resultSize, len(result.Items))
	}
}
