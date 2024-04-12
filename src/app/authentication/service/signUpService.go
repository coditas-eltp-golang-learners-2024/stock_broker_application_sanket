package service

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repo"
	"authentication/utils"
	"errors"
)

type UserRepo struct {
	userExistenceChecker repo.UserExistenceChecker
}

func NewUserRepo(checker repo.UserExistenceChecker) *UserRepo {
	return &UserRepo{
		userExistenceChecker: checker,
	}
}

// @Summary Register a new customer
// @Description Registers a new customer and saves their data in the database.
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body models.Customer true "Customer data"
// @Success 200 {string} string "success"
// @Failure 400 {object} string "Bad Request"
// @Failure 409 {object} string "Email already exists"
// @Failure 409 {object} string "Phone number already exists"
// @Failure 409 {object} string "Pan card number already exists"
// @Failure 422 {object} string "Invalid customer data"
// @Failure 500 {object} string "Internal Server Error"
func (u *UserRepo) SignUp(customer models.Customer) error {

	if u.userExistenceChecker.IsEmailExists(customer.Email) {
		return constants.ErrEmailExsits
	}
	if u.userExistenceChecker.IsPhoneNumberExists(uint64(customer.PhoneNumber)) {
		return constants.ErrInvalidPhoneNumber
	}
	if u.userExistenceChecker.IsPancardNumberExists(customer.PancardNumber) {
		return constants.ErrPanCardlExsits
	}
	if err := utils.ValidateCustomer(customer); err != nil {
		return errors.New(constants.ErrInValidation.Error() + err.Error())
	}
	if !u.userExistenceChecker.InsertCustomer(&customer) {
		return constants.ErrQueryInDB
	}
	return nil
}
