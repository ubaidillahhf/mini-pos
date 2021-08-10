package suppliers

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id          int            `json:"id" gorm:"primaryKey;index"`
	Name        string         `json:"name" gorm:"not null"`
	Address     string         `json:"address" gorm:"not null"`
	Phone       string         `json:"phone" gorm:"not null"`
	Email       string         `json:"email" gorm:"not null"`
	ProductId   int            `json:"product_id" gorm:"not null;index"`
	MerchantId  int            `json:"merchant_id" gorm:"not null;index"`
	Remark      string         `json:"remark"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
