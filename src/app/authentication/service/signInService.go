package service

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repository"
	"math/rand"
	"time"
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
	if !signInUser.credentials.CheckCredentialsExist(userCredentials.Email, userCredentials.Password) {
		return constants.ErrSignIn
	}
	otp, creationTime := generateOTP()
	if !signInUser.credentials.AssignOtpToEmail(userCredentials.Email, otp, creationTime) {
		return constants.ErrOtpGeneration
	}
	return nil
}

func generateOTP() (int, time.Time) {
	creationTime := time.Now().Add(time.Minute)
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(10000)
	return otp, creationTime
}
