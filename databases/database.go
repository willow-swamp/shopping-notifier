package databases

import (
	"github.com/willow-swamp/shopping-notifier/models"

	mysql "github.com/go-sql-driver/mysql"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConn() (*gorm.DB, error) {
	config := mysql.Config{
		User:      "user",
		Passwd:    "password",
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "gorm_database",
		ParseTime: true,
		Params: map[string]string{
			"loc": "Local",
		},
	}
	db, err := gorm.Open(gorm_mysql.Open(config.FormatDSN()), &gorm.Config{})
	return db, err
}

func Migrate() error {
	db, err := DBConn()
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.User{}, &models.Item{}, &models.Group{}, &models.Whitelist{})
	return err
}

func GetItems() ([]models.Item, error) {
	db, err := DBConn()
	if err != nil {
		return nil, err
	}
	var items []models.Item
	db.Find(&items)
	return items, nil
}

func GetItem(id int) (*models.Item, error) {
	db, err := DBConn()
	if err != nil {
		return nil, err
	}
	var item models.Item
	result := db.First(&item, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func CreateItem(item *models.Item) error {
	db, err := DBConn()
	if err != nil {
		return err
	}
	db.Create(item)
	return nil
}

func UpdateItem(item *models.Item) error {
	db, err := DBConn()
	if err != nil {
		return err
	}
	db.Updates(item)
	return nil
}

func DeleteItem(id int) error {
	db, err := DBConn()
	if err != nil {
		return err
	}
	db.Delete(&models.Item{}, id)
	return nil
}
