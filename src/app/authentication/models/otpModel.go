package models

import "time"

type User struct {
	Email       string    `gorm:"column:email;index" json:"email" validate:"required" example:"sanket@coditas.com`
	OTP        string    `gorm:"column:otp" json:"otp" validate:"required" example:"8393"`
	Otp_expiry time.Time `gorm:"column:otp_expiry" json:"otp_expiry"`
}


func (User) TableName() string {
	return "customerdata"
}
