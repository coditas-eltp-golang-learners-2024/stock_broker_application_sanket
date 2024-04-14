package repository

import (
	"authentication/models"
	"gorm.io/gorm"
)

type UserExistenceChecker interface {
	IsEmailExists(email string) bool
	IsPhoneNumberExists(phoneNumber uint64) bool
	IsPancardNumberExists(pancardNumber string) bool
	InsertCustomer(customer *models.Customer) bool
}	// SingUp

type SignInCredentials interface {
	CheckEmailExists(email string) bool
	CheckPasswordExists(password string) bool
}	//SignIn

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

func (userRepository *UserDBRepository) CheckEmailExists(email string) bool {
	var count int64
	if err := userRepository.db.Model(&models.SignInCredentials{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) CheckPasswordExists(password string) bool {
	var count int64
	if err := userRepository.db.Model(&models.SignInCredentials{}).Where("password = ?", password).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
