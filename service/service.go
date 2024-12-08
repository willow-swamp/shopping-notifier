package service

import (
	"github.com/willow-swamp/shopping-notifier/databases"
	"github.com/willow-swamp/shopping-notifier/models"
)

type Service struct {
	repository databases.Repository
}

func NewService(repository databases.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetItems() ([]models.Item, error) {
	return s.repository.GetItems()
}

func (s *Service) GetItem(id int) (*models.Item, error) {
	return s.repository.GetItem(id)
}

func (s *Service) CreateItem(item *models.Item) error {
	return s.repository.CreateItem(item)
}

func (s *Service) UpdateItem(item *models.Item) error {
	return s.repository.UpdateItem(item)
}

func (s *Service) DeleteItem(id int) error {
	return s.repository.DeleteItem(id)
}
