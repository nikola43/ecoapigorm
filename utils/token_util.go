package utils

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models"
	"math"
	"time"
)

func GenerateClientToken(email string, client_id uint, clinic_id uint) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["client_id"] = client_id
	claims["clinic_id"] = clinic_id
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

func GetClientTokenClaims(context *fiber.Ctx) (models.ClientTokenClaims, error) {
	// todo handle error
	user := context.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	clientTokenClaim := models.ClientTokenClaims{
		Email:    claims["email"].(string),
		ClientID: uint(math.Round(claims["client_id"].(float64))),
		ClinicID: uint(math.Round(claims["clinic_id"].(float64))),
		Exp:      uint(math.Round(claims["exp"].(float64))),
	}
	return clientTokenClaim, nil
}

func GetEmployeeTokenClaims(context *fiber.Ctx) (models.EmployeeTokenClaims, error) {
	// todo handle error
	user := context.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	employeeTokenClaims := models.EmployeeTokenClaims{
		Email:      claims["email"].(string),
		EmployeeID: uint(math.Round(claims["employee_id"].(float64))),
		ClinicID:   uint(math.Round(claims["clinic_id"].(float64))),
		Role:       claims["role"].(string),
		Exp:        uint(math.Round(claims["exp"].(float64))),
	}
	return employeeTokenClaims, nil
}

func GenerateEmployeeToken(email string, employee_id uint, clinic_id uint, role string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["employee_id"] = employee_id
	claims["clinic_id"] = clinic_id
	claims["email"] = email
	claims["role"] = role
	// todo añaair role e user id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte(GetEnvVariable("JWT_CLIENT_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, err
}
