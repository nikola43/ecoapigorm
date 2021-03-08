package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/utils"
)

func XApiKeyMiddleware(context *fiber.Ctx) error {
	// todo investigar por que no llega X_API_KEY
	/*
	requestApiKey := context.Get("X_API_KEY")
	serverApiKey := utils.GetEnvVariable("X_API_KEY")
	fmt.Println("requestApiKey")
	fmt.Println(requestApiKey)
	fmt.Println("serverApiKey")
	fmt.Println(serverApiKey)
	// context.h

	if requestApiKey != serverApiKey {
		return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "unauthorized",
		})
	}
	*/

	return context.Next()
}

func AdminEmployeeMiddleware(context *fiber.Ctx) error {
	if employeeTokenClaims, err := utils.GetEmployeeTokenClaims(context); err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "unauthorized",
		})
	} else if employeeTokenClaims.Role != models.ADMIN_ROLE {
		return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "employee not is admin",
		})
	} else {
		return context.Next()
	}
}
