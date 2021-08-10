package payments

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id          int            `json:"id" gorm:"primaryKey;index"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	Status      string         `json:"status" gorm:"not null"`
	Note        string         `json:"note"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
