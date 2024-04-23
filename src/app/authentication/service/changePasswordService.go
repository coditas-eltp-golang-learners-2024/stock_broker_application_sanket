package service

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repository"
	"authentication/utils"
	"errors"
)

type PasswordResetter struct {
	PasswordResetter repository.AuthenticationProvider
}

func NewRestPasswordService(restPswd repository.AuthenticationProvider) *PasswordResetter {
	return &PasswordResetter{
		PasswordResetter: restPswd,
	}
}

func (service *PasswordResetter) ResetPassword(request models.ChangePassword) error {
	if !service.PasswordResetter.CheckEmailAndPassword(request.Email, request.OldPassword) {
		return constants.ErrInvalidEmailOrPassword
	}
	if err := utils.ValidateOtpChange(request); err != nil {
		return errors.New(constants.ErrInValidation.Error() + err.Error())
	}
	if !service.PasswordResetter.SetNewPassword(request.Email, request.NewPassword) {
		return constants.ErrFailedToSetNewPassword
	}
	return nil
}
