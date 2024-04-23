package models

import "github.com/golang-jwt/jwt"

type Claim struct {
	Email string `gorm:"column:email" json:"email" example:"john.doe@gmail.com"`
	jwt.StandardClaims
}
