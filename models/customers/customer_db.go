package customers

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id                int            `json:"id" gorm:"primaryKey;index"`
	Name              string         `json:"name" gorm:"not null"`
	Address           string         `json:"address" gorm:"not null"`
	Phone             string         `json:"phone" gorm:"not null"`
	WhatsappAvailable int8           `json:"whatsapp_available" gorm:"not null"`
	Remark            string         `json:"remark"`
	Description       string         `json:"description"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
