package orderdetails

import (
	"time"

	"gorm.io/gorm"
)

type OrderDetail struct {
	Id        int            `json:"id" gorm:"primaryKey;index"`
	OrderId   int            `json:"order_id" gorm:"not null;index"`
	ProductId int            `json:"product_id" gorm:"not null;index"`
	Quantity  int            `json:"quantity" gorm:"not null"`
	Note      string         `json:"note"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
