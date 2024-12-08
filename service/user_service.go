package service

import (
	"github.com/willow-swamp/shopping-notifier/databases"
	"github.com/willow-swamp/shopping-notifier/models"
)

type UserService struct {
	repository databases.UserRepository
}

func NewUserService(repository databases.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.repository.GetUsers()
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	return s.repository.GetUser(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repository.CreateUser(user)
}
