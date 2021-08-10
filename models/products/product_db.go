package products

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id          int            `json:"id" gorm:"primaryKey;index"`
	MerchantId  int            `json:"merchant_id" gorm:"not null;index"`
	Sku         string         `json:"sku" gorm:"not null"`
	Name        string         `json:"name" gorm:"not null"`
	Remark      string         `json:"remark"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
