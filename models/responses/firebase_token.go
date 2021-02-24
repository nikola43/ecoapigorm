package models

import (
	"github.com/dgrijalva/jwt-go"
)

type FirebaseToken struct {
	Username string
	jwt.StandardClaims
}
