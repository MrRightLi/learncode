package persist

import (
	"context"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver Got item: #%d: %v", itemCount, item)
			itemCount++

			save(item)
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		// Must turn off sniffer
		elastic.SetSniff(false))

	if err != nil {
		return "", err
	}
	response, err := client.Index().
		Index("dataing_profile").
		Type("zhengai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", err
	}
	return response.Id, nil
}