package service

import (
	"github.com/willow-swamp/shopping-notifier/databases"
	"github.com/willow-swamp/shopping-notifier/models"
)

type ItemService struct {
	repository databases.ItemRepository
}

func NewItemService(repository databases.ItemRepository) *ItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) GetItems() ([]models.Item, error) {
	return s.repository.GetItems()
}

func (s *ItemService) GetItem(id int) (*models.Item, error) {
	return s.repository.GetItem(id)
}

func (s *ItemService) CreateItem(item *models.Item) error {
	return s.repository.CreateItem(item)
}

func (s *ItemService) UpdateItem(item *models.Item) error {
	return s.repository.UpdateItem(item)
}

func (s *ItemService) DeleteItem(id int) error {
	return s.repository.DeleteItem(id)
}
