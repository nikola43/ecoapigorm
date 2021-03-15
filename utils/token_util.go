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

func GenerateClientToken(email string, client_id uint, clinic_id uint) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = client_id
	claims["clinic_id"] = clinic_id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte(GetEnvVariable("JWT_CLIENT_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

// todo validar claims por separado
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

		if claims["clinic_id"] != nil {
			clientTokenClaims.ClinicID = uint(math.Round(claims["clinic_id"].(float64)))
		}

		if claims["exp"] != nil {
			clientTokenClaims.ClinicID = uint(math.Round(claims["exp"].(float64)))
		}

		return clientTokenClaims, nil
	} else {
		return nil, errors.New("invalid claims")
	}
	return nil, errors.New("invalid claims")
}

func GetEmployeeTokenClaims(context *fiber.Ctx) (*models.EmployeeTokenClaims, error) {
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

		if claims["company_id"] != nil {
			employeeTokenClaims.ClinicID = uint(math.Round(claims["company_id"].(float64)))
		}

		if claims["company_name"] != nil {
			employeeTokenClaims.CompanyName = claims["company_name"].(string)
		}

		if claims["clinic_name"] != nil {
			employeeTokenClaims.ClinicName = claims["clinic_name"].(string)
		}

		if claims["clinic_id"] != nil {
			employeeTokenClaims.ClinicID = uint(math.Round(claims["clinic_id"].(float64)))
		}

		if claims["exp"] != nil {
			employeeTokenClaims.ClinicID = uint(math.Round(claims["exp"].(float64)))
		}

		return employeeTokenClaims, nil
	} else {
		return nil, errors.New("invalid claims")
	}
	return nil, errors.New("invalid claims")
}

func GenerateEmployeeToken(name string, email string, companyName string, clinicName string, employee_id uint, company_id uint, clinic_id uint, role string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = employee_id
	claims["clinic_id"] = clinic_id
	claims["company_id"] = company_id
	claims["clinic_name"] = clinicName
	claims["company_name"] = companyName
	claims["email"] = email
	claims["name"] = name
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	fmt.Println(employee_id)
	fmt.Println(clinic_id)
	fmt.Println(company_id)
	fmt.Println(clinicName)
	fmt.Println(companyName)
	fmt.Println(employee_id)
	fmt.Println(email)
	fmt.Println(name)
	fmt.Println(role)
	fmt.Println(claims["exp"])

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte(GetEnvVariable("JWT_CLIENT_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func GenerateInvitationToken(fromEmail string, toEmail string, clinicID uint) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["from_email"] = fromEmail
	claims["to_email"] = toEmail
	claims["clinic_id"] = clinicID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(GetEnvVariable("INVITE_TOKEN")))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func GetInviteTokenClaims(context *fiber.Ctx) (*models.InviteTokenClaims, error) {
	user := context.Locals("user").(*jwt.Token)

	if claims, ok := user.Claims.(jwt.MapClaims); ok && user.Valid {
		inviteTokenClaims := &models.InviteTokenClaims{}

		if claims["from_email"] != nil {
			inviteTokenClaims.FromEmail = claims["from_email"].(string)
		}

		if claims["to_email"] != nil {
			inviteTokenClaims.ToEmail = claims["to_email"].(string)
		}

		if claims["clinic_id"] != nil {
			inviteTokenClaims.ClinicID = uint(math.Round(claims["clinic_id"].(float64)))
		}

		if claims["exp"] != nil {
			inviteTokenClaims.Exp = uint(math.Round(claims["exp"].(float64)))
		}

		return inviteTokenClaims, nil
	} else {
		return nil, errors.New("invalid claims")
	}
	return nil, errors.New("invalid claims")
}
