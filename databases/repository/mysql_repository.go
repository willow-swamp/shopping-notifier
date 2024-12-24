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

func (db MySQLRepository) GetItems(sub string) ([]models.Item, error) {
	var items []models.Item
	result := db.db.Where("group_id IN (?)", db.db.Table("users").Select("group_id").Where("line_id = ?", sub)).Find(&items)
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

func (db MySQLRepository) GetUser(sub string) (*models.User, error) {
	var user models.User
	result := db.db.Where("line_id = ?", sub).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

//func (db MySQLRepository) CreateUser(user *models.User) error {
//	result := db.db.Create(user)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}

func (db MySQLRepository) GetGroup(id uint) (*models.Group, error) {
	var group models.Group
	result := db.db.First(&group, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &group, nil
}
