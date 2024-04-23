package service

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repository"
)

type OtpVerificationService struct {
	otpVerification repository.OtpVerification
}

func NewOtpVerificationService(otpVerificationRepo repository.OtpVerification) *OtpVerificationService {
	return &OtpVerificationService{
		otpVerification: otpVerificationRepo,
	}
}

func (otpService *OtpVerificationService) OtpVerification(otpData models.User) error {
	if !otpService.otpVerification.CheckOtp(otpData.Email, otpData.OTP) {
		return constants.ErrOtpVerification
	}
	Token, err := GenerateToken(otpData.Email)
	if err != nil {
		return constants.ErrTokenGenerationErrorMessage
	}
	if !otpService.otpVerification.AddJWTTokenToDB(Token,otpData.Email){
		return constants.ErrTokenFailedToInsertIntoDB
	}

	return nil
}
