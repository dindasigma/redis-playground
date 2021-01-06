package items

import (
	"encoding/json"
	"log"

	"github.com/dindasigma/my-playground/go-es/client/elasticsearch"
)

const (
	indexItems = "items"
)

func (i *Item) Save() error {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return err
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() error {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, i.Id)
	if err != nil {
		return err
	}
	log.Print(result.Error)

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, i); err != nil {
		return err
	}

	i.Id = itemId
	log.Print(string(bytes))
	return nil

}
