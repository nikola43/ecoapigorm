package utils

import (
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
)

type customToken struct {
	Username string
	jwt.StandardClaims
}

func GenerateTokenUsername(userName string) string {
	// Declare the expiration time of the token
	// here, we have kept it as 1 day
	expirationTime := time.Now().Add(1440 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &customToken{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// generate api token
	apiToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	apiTokenString, _ := apiToken.SignedString([]byte(os.Getenv("token_password")))
	return apiTokenString
}
