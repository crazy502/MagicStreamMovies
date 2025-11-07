package utils

import (
	jwt "github.com/golang-jwt/jwt/v5"
)

// 将被编码到相关的JWT中
type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Role      string
	UserId    string
	jwt.ResgisteredClaims
}
