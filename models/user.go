package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	LineId  string
	Name    string
	GroupId int
}
