package models

import (
	"github.com/willow-swamp/shopping-notifier/config"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	GroupID     uint                   `gorm:"not null"`
	Name        string                 `gorm:"not null"`
	Priority    config.PriorityType    `gorm:"type:enum('低', '中', '高'):default:'中'"`
	StockStatus config.StockStatusType `gorm:"type:enum('在庫あり', '在庫なし'):default:'在庫なし'"`
}
