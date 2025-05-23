package model

import (
	"time"
)

type Cart struct {
	CartID     int        `gorm:"column:cart_id;AUTO_INCREMENT;primary_key"`
	CustomerID int        `gorm:"column:customer_id;NOT NULL"`
	CartName   string     `gorm:"column:cart_name"`
	CreatedAt  time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	CartItems  []CartItem `gorm:"foreignKey:CartID"` // ใช้ Slice เพื่อรองรับหลาย CartItem
}

func (m *Cart) TableName() string {
	return "cart"
}
