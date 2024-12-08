package repository

import (
	"github.com/willow-swamp/shopping-notifier/models"

	"gorm.io/gorm"
)

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (db MySQLRepository) GetItems() ([]models.Item, error) {
	var items []models.Item
	result := db.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func (db MySQLRepository) GetItem(id int) (*models.Item, error) {
	var item models.Item
	result := db.db.First(&item, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func (db MySQLRepository) CreateItem(item *models.Item) error {
	result := db.db.Create(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db MySQLRepository) UpdateItem(item *models.Item) error {
	result := db.db.Updates(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db MySQLRepository) DeleteItem(id int) error {
	result := db.db.Delete(&models.Item{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db MySQLRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	result := db.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (db MySQLRepository) GetUser(id int) (*models.User, error) {
	var user models.User
	result := db.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (db MySQLRepository) CreateUser(user *models.User) error {
	result := db.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
