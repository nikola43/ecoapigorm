package utils

import (
	"errors"
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

// todo validar claims por separado
func GetClientTokenClaims(context *fiber.Ctx) (models.ClientTokenClaims, error) {
	user := context.Locals("user").(*jwt.Token)
	if claims, ok := user.Claims.(jwt.MapClaims); ok && user.Valid {
		if claims["email"] != nil && claims["client_id"] != nil && claims["clinic_id"] != nil && claims["exp"] != nil {
			clientTokenClaims := models.ClientTokenClaims{
				Email:    claims["email"].(string),
				ClientID: uint(math.Round(claims["client_id"].(float64))),
				ClinicID: uint(math.Round(claims["clinic_id"].(float64))),
				Exp:      uint(math.Round(claims["exp"].(float64))),
			}
			return clientTokenClaims, nil
		}

	} else {
		return models.ClientTokenClaims{}, errors.New("invalid claims")
	}
	return models.ClientTokenClaims{}, errors.New("invalid claims")
}

func GetEmployeeTokenClaims(context *fiber.Ctx) (models.EmployeeTokenClaims, error) {
	user := context.Locals("user").(*jwt.Token)
	if claims, ok := user.Claims.(jwt.MapClaims); ok && user.Valid {
		if claims["email"] != nil && claims["role"] != nil && claims["clinic_id"] != nil && claims["exp"] != nil {
			employeeTokenClaims := models.EmployeeTokenClaims{
				Email:    claims["email"].(string),
				Role:     claims["role"].(string),
				ClinicID: uint(math.Round(claims["clinic_id"].(float64))),
				Exp:      uint(math.Round(claims["exp"].(float64))),
			}
			return employeeTokenClaims, nil
		}

	} else {
		return models.EmployeeTokenClaims{}, errors.New("invalid claims")
	}
	return models.EmployeeTokenClaims{}, errors.New("invalid claims")

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
