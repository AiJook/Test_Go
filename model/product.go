package model

import (
	"time"
)

type Product struct {
	ProductID     int       `gorm:"column:product_id;AUTO_INCREMENT;primary_key"`
	ProductName   string    `gorm:"column:product_name;NOT NULL"`
	Description   string    `gorm:"column:description"`
	Price         string    `gorm:"column:price;NOT NULL"`
	StockQuantity int       `gorm:"column:stock_quantity;NOT NULL"`
	CreatedAt     time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
}

func (m *Product) TableName() string {
	return "product"
}
