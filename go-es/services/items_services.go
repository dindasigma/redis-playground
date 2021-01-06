package services

import (
	"github.com/dindasigma/my-playground/go-es/domain/items"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, error)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, error) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}
