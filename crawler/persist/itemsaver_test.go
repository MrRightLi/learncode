package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"project/learncode/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {

	excepted := model.Profile{
		Age:34,
		Height:162,
		Weight:"57",
		Income:"3001-500元",
		Gender:"女",
		Name:"安静的雪",
		Xinzuo:"牡羊座",
		Occupation:"人事",
		Marriage:"离异",
		House:"已购房",
		Hokou:"山东菏泽",
		Education:"大学本科",
		Car:"未购车",
	}

	id, err := save(excepted)

	if err != nil {
		panic(err)
	}

	// TODO: Try to start up elastic search
	// hear using docker go clint
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	response, err := client.Get().
		Index("dataing_profile").
		Type("zhengai").
		Id(id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", response.Source)

	var actual model.Profile
	err = json.Unmarshal([]byte(*response.Source), &actual)

	if err != nil {
		panic(err)
	}

	if actual != excepted {
		t.Errorf("got %v; expected: %v", actual, excepted)
	} else {
		fmt.Printf("got actual: %+v", actual)
	}
}
