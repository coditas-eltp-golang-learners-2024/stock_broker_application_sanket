package models

// import "github.com/golang-jwt/jwt/v4"

// Customer represents a customer model.
type Customer struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"column:name" json:"name" validate:"required,min=3,max=50" example:"John Doe"`
	Email         string `gorm:"column:email" json:"email" validate:"required,email" example:"john.doe@gmail.com"`
	PhoneNumber   uint64 `gorm:"column:phoneNumber" json:"phoneNumber" validate:"required,lt=10000000000,gt=999999999" example:"1234567890"`
	PancardNumber string `gorm:"column:pancard" json:"pancard" validate:"required,len=10" example:"ABCDE1234F"`
	Password      string `gorm:"column:password" json:"password" validate:"required,min=8" example:"password"`
}

func (Customer) TableName() string {
	return "users"
}
