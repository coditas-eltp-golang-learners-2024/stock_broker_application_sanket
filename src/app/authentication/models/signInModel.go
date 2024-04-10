package models

type UserCredentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

var UerCredentialsList []UserCredentials