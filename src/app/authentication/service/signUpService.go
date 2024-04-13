package service

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repository"
	"authentication/utils"
	"errors"
)

type UserRepo struct {
	userExistenceChecker repository.UserExistenceChecker
}

func NewUserRepo(checker repository.UserExistenceChecker) *UserRepo {
	return &UserRepo{
		userExistenceChecker: checker,
	}
}

func (user *UserRepo) SignUp(customer models.Customer) error {
	if user.userExistenceChecker.IsEmailExists(customer.Email) {
		return constants.ErrEmailExsits
	}
	if user.userExistenceChecker.IsPhoneNumberExists(uint64(customer.PhoneNumber)) {
		return constants.ErrInvalidPhoneNumber
	}
	if user.userExistenceChecker.IsPancardNumberExists(customer.PancardNumber) {
		return constants.ErrPanCardlExsits
	}
	if err := utils.ValidateCustomer(customer); err != nil {
		return errors.New(constants.ErrInValidation.Error() + err.Error())
	}
	if !user.userExistenceChecker.InsertCustomer(&customer) {
		return constants.ErrQueryInDB
	}
	return nil
}
