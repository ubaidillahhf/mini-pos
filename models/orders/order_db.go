package orders

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id         int            `json:"id" gorm:"primaryKey;index"`
	UserId     int            `json:"user_id" gorm:"not null;index"`
	CustomerId int            `json:"customer_id" gorm:"not null;index"`
	OutletId   int            `json:"outlet_id" gorm:"not null;index"`
	PaymentId  int            `json:"payment_id" gorm:"not null;index"`
	Note       string         `json:"note"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
