package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	LineID  string `gorm:"unique;not null"`
	GroupID int    `gorm:"not null"`
}
