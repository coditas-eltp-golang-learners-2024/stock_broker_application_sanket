package models

import "time"

type User struct {
	Email        string    `gorm:"column:email;index" json:"email" validate:"required" example:"sanket@gmail.com"`
	OTP          int       `gorm:"column:otp" json:"otp" validate:"required" example:"8393"`
	CreationTime time.Time `gorm:"column:createdAt" json:"createdAt" example:"2024-04-17 07:39:20"`
}

func (User) TableName() string {
	return "users"
}
