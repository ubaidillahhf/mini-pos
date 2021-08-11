package orderdetails

import (
	"time"

	"gorm.io/gorm"
)

type Merchant struct {
	Id               int            `json:"id" gorm:"primaryKey;index"`
	Name             string         `json:"name" gorm:"not null"`
	UserId           int            `json:"user_id" gorm:"not null;index"`
	Npwp             string         `json:"npwp" gorm:"not null"`
	BusinnessAddress string         `json:"businness_address" gorm:"not null"`
	BusinnessPhone   string         `json:"businness_phone" gorm:"not null"`
	RegisterAddress  string         `json:"register_address" gorm:"not null"`
	RegisterPhone    string         `json:"register_phone" gorm:"not null"`
	DisplayImage     string         `json:"display_image" gorm:"not null"`
	Note             string         `json:"note"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
