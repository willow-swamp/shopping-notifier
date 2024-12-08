package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupName string `gorm:"not null"`
	Users     []User `gorm:"foreignKey:GroupID"`
	Items     []Item `gorm:"foreignKey:GroupID"`
}
