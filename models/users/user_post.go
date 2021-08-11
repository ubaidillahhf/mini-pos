package users

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreate struct {
	Name     string `json:"name" validate:"required,min=2"`
	Address  string `json:"address" validate:"required"`
	Phone    string `json:"phone" validate:"required" gorm:"not null"`
	Email    string `json:"email" validate:"required,email" gorm:"not null" `
	Password string `json:"password" validate:"required" gorm:"not null"`
	Remark   string `json:"remark"`
}

type UserEdit struct {
	Name            string `json:"name" validate:"required,min=2"`
	Address         string `json:"address" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	NewPassword     string `json:"newPassword" validate:"required"`
	RetypePassword  string `json:"retypePassword"`
	ConfirmPassword string `json:"confirmPassword"`
}
