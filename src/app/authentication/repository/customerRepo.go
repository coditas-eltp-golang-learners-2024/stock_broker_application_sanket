package repository

import (
	"authentication/models"
	"time"

	"gorm.io/gorm"
)

type UserExistenceChecker interface {
	IsEmailExists(email string) bool
	IsPhoneNumberExists(phoneNumber uint64) bool
	IsPancardNumberExists(pancardNumber string) bool
	InsertCustomer(customer *models.Customer) bool
} // SingUp

type SignInCredentials interface {
	CheckPasswordExists(email, password string) bool
} //SignIn

type OtpVerification interface {
	CheckEmailExists(email string) bool
	NewRecordInsert(email, otp string, otp_expiry time.Time) bool
	CheckOtp(email string, otp string) bool
} //OtpVerification

type UserDBRepository struct {
	db *gorm.DB
}

func NewUserDBRepository(dataBase *gorm.DB) *UserDBRepository {
	return &UserDBRepository{db: dataBase}
}

func (userRepository *UserDBRepository) IsEmailExists(email string) bool {
	var count int64
	if err := userRepository.db.Model(&models.Customer{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) IsPhoneNumberExists(phoneNumber uint64) bool {
	var count int64
	if err := userRepository.db.Model(&models.Customer{}).Where("phone_number = ?", phoneNumber).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) IsPancardNumberExists(pancardNumber string) bool {
	var count int64
	if err := userRepository.db.Model(&models.Customer{}).Where("pancard_number = ?", pancardNumber).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) InsertCustomer(customerData *models.Customer) bool {
	// Create the customer record in the database
	if err := userRepository.db.Create(&customerData).Error; err != nil {
		return false
	}
	return true
}

func (userRepository *UserDBRepository) CheckPasswordExists(email, password string) bool {
	var count int64
	if err := userRepository.db.Model(&models.SignInCredentials{}).
		Where("email = ? AND password = ?", email, password).
		Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) CheckEmailExists(email string) bool {
	var count int64
	if err := userRepository.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) NewRecordInsert(email, otp string, otpExpiry time.Time) bool {
	var count int64

	// Truncate milliseconds from otpExpiry
	otpExpiry = otpExpiry.Truncate(time.Second)
	currentTime := time.Now()
	if otpExpiry.Sub(currentTime) <= time.Minute {
		if err := userRepository.db.Model(&models.User{}).Where("email = ?", email).Updates(models.User{OTP: otp, Otp_expiry: otpExpiry}).Count(&count).Error; err != nil {
			return false
		}
	}

	// Update existing record with new OTP and OTP expiry
	// if err := userRepository.db.Model(&models.User{}).Where("email = ?", email).Updates(models.User{OTP: otp, Otp_expiry: otpExpiry}).Count(&count).Error; err != nil {
	// 	return false
	// }
	return count > 0
}

func (userRepository *UserDBRepository) CheckOtp(email, otp string) bool {
	var count int64
	if err := userRepository.db.Model(&models.User{}).Where("otp = AND email?", otp, email).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
