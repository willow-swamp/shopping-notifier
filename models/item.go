package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	GroupID     int
	Name        string
	Priority    int
	StockStatus string
}
