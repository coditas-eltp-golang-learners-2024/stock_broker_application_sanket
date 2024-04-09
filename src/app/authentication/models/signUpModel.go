package models

type Customer struct {
	ID            int    `json:"id"`
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required"`
	PhoneNumber   string `json:"phoneNumber" validate:"required,len=10"`
	PanCardNumber string `json:"panCardNumber" validate:"required"`
	Password      string `json:"password" validate:"required"`
}

var CustomerData []Customer
