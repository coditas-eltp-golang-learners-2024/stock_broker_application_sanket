package models

import "time"

type User struct {
	Email        string    `gorm:"column:email;index" json:"email" validate:"required" example:"sanket@gmail.com`
	OTP          string    `gorm:"column:otp" json:"otp" validate:"required" example:"8393"`
	CreationTime time.Time `gorm:"column:creation_Time" json:"creationTime"`
}

func (User) TableName() string {
	return "customerdata"
}
