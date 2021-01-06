package items

import (
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
