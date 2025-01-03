package databases

import (
	"github.com/willow-swamp/shopping-notifier/models"
)

type ItemRepository interface {
	GetItems(sub string) ([]models.Item, error)
	GetItem(id int) (*models.Item, error)
	CreateItem(item *models.Item) error
	UpdateItem(item *models.Item) error
	DeleteItem(id int) error
}

type UserRepository interface {
	GetUsers() ([]models.User, error)
	GetUser(sub string) (*models.User, error)
	//CreateUser(user *models.User) error
}

type GroupRepository interface {
	GetGroup(id uint) (*models.Group, error)
}
