package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	GroupID     int    `gorm:"not null"`
	Name        string `gorm:"not null"`
	Priority    int    `gorm:"not null;default:2"`
	StockStatus int    `gorm:"not null;default:2"`
}
