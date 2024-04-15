package service

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repository"
	"crypto/rand"
	"math/big"
	"time"
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

	if !otpService.otpVerification.CheckEmailExists(otpData.Email) {
		return constants.ErrInvalidName
	}
	otp,otp_expiry:=generateOTP()
	if otpService.otpVerification.NewRecordInsert(otpData.Email,otp,otp_expiry) {
		return constants.ErrQueryInDB
	}
	if otpService.otpVerification.CheckOtp(otpData.Email,otpData.OTP) {
		return constants.ErrCheckingPanCard
	}
	return nil
}

func generateOTP() (string, time.Time) {
	expiry := time.Now().Add(time.Minute * constants.ExpiryMins)
	otp := ""
	for i := 0; i < constants.OtpLength; i++ {
		randomNumber, _ := rand.Int(rand.Reader, big.NewInt(10))
		otp += randomNumber.String()
	}
	return otp, expiry
}
