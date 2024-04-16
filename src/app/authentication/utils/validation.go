package utils

import (
	"authentication/constants"
	"authentication/models"
	"errors"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate
func init() {
	Validate = validator.New()
}

func ValidateCustomer(customer models.Customer) error {
	err := Validate.Struct(customer)
	if err != nil {
		validationErrors, _ := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			structFieldName := fieldError.Field()
			switch structFieldName {
			case "Name":
				return errors.New(constants.ErrInvalidName.Error())
			case "Email":
				return errors.New(constants.ErrInvalidEmail.Error())
			case "PhoneNumber":
				return errors.New(constants.ErrInvalidPhoneNumber.Error())
			case "PancardNumber":
				return errors.New(constants.ErrInvalidPanCard.Error())
			case "Password":
				return errors.New(constants.ErrInvalidPassword.Error())
			}
		}
		return err 
	}
	return nil 
}

func ValidateSignIn(customer models.SignInCredentials) error {
	err := Validate.Struct(customer)
	if err != nil {
		validationErrors, _ := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			structFieldName := fieldError.Field()
			switch structFieldName {
			case "Email":
				return errors.New(constants.ErrInvalidEmail.Error())
			case "Password":
				return errors.New(constants.ErrInvalidPassword.Error())
			}
		}
		return err 
	}
	return nil 
}
