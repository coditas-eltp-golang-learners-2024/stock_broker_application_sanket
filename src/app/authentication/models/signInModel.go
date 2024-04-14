package models

type SignInCredentials struct {
	Email    string `gorm:"column:email" json:"email" validate:"required,email,contains=@coditas.com" example:"john.doe@coditas.com"`
	Password string `gorm:"column:password" json:"password" validate:"required,min=8" example:"password"`
}
