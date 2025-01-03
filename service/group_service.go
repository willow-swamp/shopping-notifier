package service

import (
	"github.com/willow-swamp/shopping-notifier/databases"
	"github.com/willow-swamp/shopping-notifier/models"
)

type GroupService struct {
	repository databases.GroupRepository
}

func NewGroupService(repository databases.GroupRepository) *GroupService {
	return &GroupService{repository: repository}
}

func (s *GroupService) GetGroup(id uint) (*models.Group, error) {
	return s.repository.GetGroup(id)
}
