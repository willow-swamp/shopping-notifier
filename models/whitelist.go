package models

import "gorm.io/gorm"

type Whitelist struct {
	gorm.Model
	LineId string
}
