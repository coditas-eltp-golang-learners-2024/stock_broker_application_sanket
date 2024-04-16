package service

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repository"
	"authentication/utils"
	"errors"
)

type SignInChecker struct {
	credentials repository.SignInCredentials
}

func NewSignInChecker(check repository.SignInCredentials) *SignInChecker {
	return &SignInChecker{
		credentials: check,
	}
}

func (signInUser *SignInChecker) SignIn(userCredentials models.SignInCredentials) error {
	if err := utils.ValidateSignIn(userCredentials); err != nil {
		return errors.New(constants.ErrInValidation.Error() + err.Error())
	}
	if !signInUser.credentials.CheckCredentialsExist(userCredentials.Email, userCredentials.Password) {
		return constants.ErrSignIn
	}
	return nil
}
