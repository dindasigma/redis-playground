package items

import (
	"encoding/json"
	"log"

	"github.com/dindasigma/my-playground/go-es/client/elasticsearch"
	"github.com/dindasigma/my-playground/go-es/domain/queries"
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

func (i *Item) Search(query queries.EsQuery) ([]Item, error) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, err
	}

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, err
		}
		item.Id = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, err
	}
	return items, nil
}
