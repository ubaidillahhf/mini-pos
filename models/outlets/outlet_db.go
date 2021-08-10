package outlets

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id         int            `json:"id" gorm:"primaryKey;index"`
	Name       string         `json:"name" gorm:"not null"`
	Address    string         `json:"address" gorm:"not null"`
	PicName    string         `json:"pic_name" gorm:"not null"`
	PicPhone   string         `json:"pic_phone" gorm:"not null"`
	MerchantId int            `json:"merchant_id" gorm:"not null;index"`
	Note       string         `json:"note"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
