package productprices

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int            `json:"id" gorm:"primaryKey;index"`
	ProductId int            `json:"product_id" gorm:"not null;index"`
	Price     float64        `json:"total_price" gorm:"not null"`
	OutletId  int            `json:"outlet_id" gorm:"not null;index"`
	Note      string         `json:"note"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
