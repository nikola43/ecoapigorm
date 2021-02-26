package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type customToken struct {
	Username string
	jwt.StandardClaims
}

func GenerateClientToken(email string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	// todo añaair role e user id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte(GetEnvVariable("JWT_CLIENT_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, err
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
