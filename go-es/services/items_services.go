package services

import (
	"github.com/dindasigma/my-playground/go-es/domain/items"
	"github.com/dindasigma/my-playground/go-es/domain/queries"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, error)
	Get(string) (*items.Item, error)
	Search(query queries.EsQuery) ([]items.Item, error)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, error) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, error) {
	item := items.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil

}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, error) {
	dao := items.Item{}
	return dao.Search(query)
}
