package models

type SignInCredentials struct {
	Email    string `gorm:"column:email" json:"email" validate:"required,email" example:"john.doe@gmail.com"`
	Password string `gorm:"column:password" json:"password" validate:"required,min=8" example:"password"`
}

func (SignInCredentials) TableName() string {
	return "customerdata"
}
