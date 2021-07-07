package utils

import (
	"errors"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models"
	"math"
	"time"
)

func GenerateClientToken(email string, clientId uint) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = clientId
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	fmt.Println(claims["id"])
	fmt.Println(claims["email"])
	fmt.Println(claims["exp"])

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte(GetEnvVariable("JWT_CLIENT_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetClientTokenClaims(context *fiber.Ctx) (*models.ClientTokenClaims, error) {
	user := context.Locals("user").(*jwt.Token)
	if claims, ok := user.Claims.(jwt.MapClaims); ok && user.Valid {
		clientTokenClaims := &models.ClientTokenClaims{}

		if claims["id"] != nil {
			clientTokenClaims.ID = uint(math.Round(claims["id"].(float64)))
		}

		if claims["email"] != nil {
			clientTokenClaims.Email = claims["email"].(string)
		}

		if claims["exp"] != nil {
			clientTokenClaims.Exp = uint(math.Round(claims["exp"].(float64)))
		}

		return clientTokenClaims, nil
	} else {
		return nil, errors.New("invalid claims")
	}
}

func GetEmployeeTokenClaims(context *fiber.Ctx) (*models.EmployeeTokenClaims, error) {
	if context.Locals("user") == nil {
		return nil, errors.New("invalid token")
	}

	user := context.Locals("user").(*jwt.Token)
	if claims, ok := user.Claims.(jwt.MapClaims); ok && user.Valid {
		employeeTokenClaims := &models.EmployeeTokenClaims{}

		if claims["id"] != nil {
			employeeTokenClaims.ID = uint(math.Round(claims["id"].(float64)))
		}

		if claims["email"] != nil {
			employeeTokenClaims.Email = claims["email"].(string)
		}

		if claims["name"] != nil {
			employeeTokenClaims.Name = claims["name"].(string)
		}

		if claims["role"] != nil {
			employeeTokenClaims.Role = claims["role"].(string)
		}

		if claims["clinic_id"] != nil {
			employeeTokenClaims.ClinicID = uint(math.Round(claims["clinic_id"].(float64)))
		}

		if claims["clinic_name"] != nil {
			employeeTokenClaims.ClinicName = claims["clinic_name"].(string)
		}

		if claims["company_id"] != nil {
			employeeTokenClaims.CompanyID = uint(math.Round(claims["company_id"].(float64)))
		}

		if claims["company_name"] != nil {
			employeeTokenClaims.CompanyName = claims["company_name"].(string)
		}

		if claims["exp"] != nil {
			employeeTokenClaims.Exp = uint(math.Round(claims["exp"].(float64)))
		}

		return employeeTokenClaims, nil
	}
	return nil, errors.New("invalid claims")
}

func GenerateEmployeeToken(name string,
	companyId uint,
	clinicId uint,
	employeeId uint,
	email string,
	companyName string,
	role string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = employeeId
	claims["clinic_id"] = clinicId
	claims["company_id"] = companyId
	claims["company_name"] = companyName
	claims["email"] = email
	claims["name"] = name
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	fmt.Println("token claims")
	fmt.Println(employeeId)
	fmt.Println(clinicId)
	fmt.Println(companyId)
	fmt.Println(companyName)
	fmt.Println(email)
	fmt.Println(name)
	fmt.Println(role)
	fmt.Println(claims["exp"])
	fmt.Println("token claims -----------")
	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte(GetEnvVariable("JWT_EMPLOYEE_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateInvitationToken() (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(GetEnvVariable("INVITE_TOKEN")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GeneratePasswordRecoveryToken(recoveryType string, id uint) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["type"] = recoveryType
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(GetEnvVariable("INVITE_TOKEN")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
