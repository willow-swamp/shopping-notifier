package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	LineID  string `gorm:"unique;not null"`
	Name    string `gorm:"not null"`
	GroupID uint   `gorm:"not null"`
}
