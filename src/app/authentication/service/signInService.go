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
	otp, creationTime := generateOTP()
	if !signInUser.credentials.AssignOtpToEmail(userCredentials.Email, otp, creationTime) {
		return constants.ErrOtpGeneration
	}
	return nil
}

func generateOTP() (string, time.Time) {
	creationTime := time.Now().Add(time.Minute * constants.ExpiryMins)
	otp := ""
	for digitIndex := 0; digitIndex < constants.OtpLength; digitIndex++ {
		randomNumber, _ := rand.Int(rand.Reader, big.NewInt(10))
		otp += randomNumber.String()
	}
	return otp, creationTime
}
