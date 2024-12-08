package databases

import (
	"github.com/willow-swamp/shopping-notifier/models"
)

type Repository interface {
	GetItems() ([]models.Item, error)
	GetItem(id int) (*models.Item, error)
	CreateItem(item *models.Item) error
	UpdateItem(item *models.Item) error
	DeleteItem(id int) error
}
