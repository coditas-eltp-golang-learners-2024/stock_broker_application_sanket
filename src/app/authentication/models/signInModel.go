package models

type SignInCredentials struct {
	Email    string `gorm:"column:email" json:"email" example:"john.doe@coditas.com"`
	Password string `gorm:"column:password" json:"password" example:"password"`
}

func (SignInCredentials) TableName() string {
	return "customerdata"
}
