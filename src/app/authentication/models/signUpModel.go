package models

// Customer represents a customer model.
type Customer struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"column:name;index" json:"name" validate:"required,min=3,max=50" example:"John Doe"`
	Email         string `gorm:"column:email;uniqueIndex" json:"email" validate:"required,email,contains=@coditas.com" example:"john.doe@coditas.com"`
	PhoneNumber   uint64 `gorm:"column:phone_number;uniqueIndex" json:"phone_number" validate:"required,lt=10000000000,gt=999999999" example:"1234567890"`
	PancardNumber string `gorm:"column:pancard_number;uniqueIndex" json:"pancard_number" validate:"required,len=10" example:"ABCDE1234F"`
	Password      string `gorm:"column:password" json:"password" validate:"required,min=8" example:"password"`
}

func (Customer) TableName() string {
	return "customerdata"
}
