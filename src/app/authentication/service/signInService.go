package service

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repository"
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
	if !signInUser.credentials.CheckPasswordExists(userCredentials.Email,userCredentials.Password) {
		return constants.ErrSignIn
	}
	return nil
}
