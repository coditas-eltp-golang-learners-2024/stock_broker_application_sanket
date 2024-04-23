package repository

import (
	"authentication/models"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

type UserExistenceChecker interface {
	IsEmailExists(email string) bool
	IsPhoneNumberExists(phoneNumber uint64) bool
	IsPancardNumberExists(pancardNumber string) bool
	InsertCustomer(customer *models.Customer) bool
} // SingUp

type SignInCredentials interface {
	CheckCredentialsExist(email, password string) bool
	AssignOtpToEmail(email string, otp int, creatime time.Time) bool
} //SignIn

type OtpVerification interface {
	CheckOtp(email string, otp int) bool
	AddJWTTokenToDB(token, email string) bool
} //OtpVerification

type AuthenticationProvider interface {
	CheckEmailAndPassword(email, password string) bool
	SetNewPassword(email, newPassword string) bool
}
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
	if err := userRepository.db.Model(&models.Customer{}).Where("phoneNumber = ?", phoneNumber).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) IsPancardNumberExists(pancardNumber string) bool {
	var count int64
	if err := userRepository.db.Model(&models.Customer{}).Where("pancard = ?", pancardNumber).Count(&count).Error; err != nil {
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

func (userRepository *UserDBRepository) CheckCredentialsExist(email, password string) bool {
	var count int64
	if err := userRepository.db.Model(&models.SignInCredentials{}).
		Where("email = ? AND password = ?", email, password).
		Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) AssignOtpToEmail(email string, otp int, creationTime time.Time) bool {
	var count int64
	// Truncate milliseconds from otpExpiry
	creationTime = creationTime.Truncate(time.Second)
	if err := userRepository.db.Model(&models.User{}).Where("email = ?", email).Updates(models.User{OTP: otp, CreationTime: creationTime}).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) CheckOtp(email string, otp int) bool {
	var count int64
	var otpCreationTime mysql.NullTime
	err := userRepository.db.Table("users").Select("createdAt").Where("email = ?", email).Scan(&otpCreationTime).Error
	if err != nil {
		return false
	}
	if otpCreationTime.Valid {
		duration := time.Since(otpCreationTime.Time)
		if duration > 1*time.Minute {
			return false
		}
	}
	if err := userRepository.db.Model(&models.User{}).Where("email=? AND otp =?", email, otp).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) AddJWTTokenToDB(token, email string) bool {
	var count int64
	if err := userRepository.db.Model(&models.Claim{}).Where("email = ?", email).Updates(models.Claim{Email: email, Token: token}).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) CheckEmailAndPassword(email, password string) bool {
	var count int64
	if err := userRepository.db.Model(&models.ChangePassword{}).Where("email = ? AND password = ?", email, password).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) SetNewPassword(email, newPassword string) bool {
	var count int64
	if err := userRepository.db.Model(&models.ChangePassword{}).Where("email = ?", email).Update("password", newPassword).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
