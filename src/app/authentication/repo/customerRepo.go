package repo

import (
	"authentication/models"
	"gorm.io/gorm"
)

type UserExistenceChecker interface {
	IsEmailExists(email string) bool
	IsPhoneNumberExists(phoneNumber uint64) bool
	IsPancardNumberExists(pancardNumber string) bool
	InsertCustomer(customer *models.Customer) bool
}

type UserDBRepository struct {
	db *gorm.DB
}

func NewUserDBRepository(Db *gorm.DB) *UserDBRepository {
	return &UserDBRepository{db: Db}
}

// IsEmailExists checks if the given email exists in the database.
// @Summary Check if email exists
// @Description Checks if the provided email address exists in the database.
// @Tags customers
// @Accept json
// @Produce json
// @Param email query string true "Email address to check"
// @Success 200 {boolean} boolean "true if email exists, false otherwise"
// @Failure 500 {object} string "Internal Server Error"
func (r *UserDBRepository) IsEmailExists(email string) bool {
	var count int64
	if err := r.db.Model(&models.Customer{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

// IsPhoneNumberExists checks if the given phone number exists in the database.
// @Summary Check if phone number exists
// @Description Checks if the provided phone number exists in the database.
// @Param phoneNumber query uint64 true "Phone number to check"
// @Success 200 {boolean} boolean "true if phone number exists, false otherwise"
// @Failure 500 {object} string "Internal Server Error"
func (r *UserDBRepository) IsPhoneNumberExists(phoneNumber uint64) bool {
	var count int64
	if err := r.db.Model(&models.Customer{}).Where("phone_number = ?", phoneNumber).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

// IsPancardNumberExists checks if the given PAN card number exists in the database.
// @Summary Check if PAN card number exists
// @Description Checks if the provided PAN card number exists in the database.
// @Param pancardNumber query string true "PAN card number to check"
// @Success 200 {boolean} boolean "true if PAN card number exists, false otherwise"
// @Failure 500 {object} string "Internal Server Error"
func (r *UserDBRepository) IsPancardNumberExists(pancardNumber string) bool {
	var count int64
	if err := r.db.Model(&models.Customer{}).Where("pancard_number = ?", pancardNumber).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

// InsertCustomer inserts a new customer record into the database.
// @Summary Insert new customer
// @Description Inserts a new customer record into the database.
// @Tags customer-signup
// @Accept json
// @Produce json
// @Param customerData body models.Customer true "Customer data"
// @Success 200 {boolean} boolean "true if customer inserted successfully, false otherwise"
// @Failure 500 {object} string "Internal Server Error"
func (r *UserDBRepository) InsertCustomer(customerData *models.Customer) bool {
	// Create the customer record in the database
	if err := r.db.Create(&customerData).Error; err != nil {
		return false
	}
	return true
}