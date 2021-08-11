package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int            `json:"id" gorm:"primaryKey;index"`
	Name      string         `json:"name" gorm:"not null"`
	Address   string         `json:"address" gorm:"not null"`
	Phone     string         `json:"phone" gorm:"not null"`
	Email     string         `json:"email" gorm:"not null;unique"`
	Password  string         `json:"password" gorm:"not null"`
	Remark    string         `json:"remark"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
