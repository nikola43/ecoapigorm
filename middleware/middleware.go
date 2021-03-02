package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/utils"
)

func ApiKeyMiddleware(context *fiber.Ctx) error {
	requestApiKey := context.Get("x-api-key")
	serverApiKey := utils.GetEnvVariable("X_API_KEY")

	if requestApiKey != serverApiKey {
		return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "unauthorized",
		})
	}

	return context.Next()
}

func AdminEmployeeMiddleware(context *fiber.Ctx) error {
	var employeeTokenClaims = models.EmployeeTokenClaims{}
	var err error

	employeeTokenClaims, err = utils.GetEmployeeTokenClaims(context)
	fmt.Println(employeeTokenClaims)
	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "unauthorized",
		})
	}
	fmt.Println(employeeTokenClaims)

	if employeeTokenClaims.Role != "admin" {
		return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "employee not is admin",
		})
	}

	return context.Next()
}
