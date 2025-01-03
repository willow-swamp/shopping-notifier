package databases

import (
	"os"

	mysql "github.com/go-sql-driver/mysql"
	"github.com/willow-swamp/shopping-notifier/models"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConn() (*gorm.DB, error) {
	config := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		ParseTime:            true,
		AllowNativePasswords: true,
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
