package config

type StockStatusType string

const (
	StockStatusInStock    StockStatusType = "在庫あり" // 在庫あり
	StockStatusOutOfStock StockStatusType = "在庫なし" // 在庫なし
)

type PriorityType string

const (
	PriorityHigh   PriorityType = "高" // 高
	PriorityMedium PriorityType = "中" // 中
	PriorityLow    PriorityType = "低" // 低
)
