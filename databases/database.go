package databases

import (
	mysql "github.com/go-sql-driver/mysql"
	"github.com/willow-swamp/shopping-notifier/models"
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

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Group{}, &models.Item{}, &models.User{}, &models.Whitelist{})
	return err
}
