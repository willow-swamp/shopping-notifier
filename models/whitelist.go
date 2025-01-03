package models

import "gorm.io/gorm"

type Whitelist struct {
	gorm.Model
	LineID string `gorm:"unique;not null"`
}
