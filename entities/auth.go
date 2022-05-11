package entities

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Auth struct {
	Token string
}

type Claims struct {
	Username string
	ID       uint64
	gorm.Model
	jwt.StandardClaims
}
