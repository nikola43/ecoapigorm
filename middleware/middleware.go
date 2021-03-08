package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/utils"
)

func XApiKeyMiddleware(context *fiber.Ctx) error {
	requestApiKey := context.Get("X_API_KEY")
	// serverApiKey := utils.GetEnvVariable("X_API_KEY")
	// todo investigar por que no coge .env
	serverApiKey := "ef2ff59e253e5c36f7f11a387c2c4a1c33ed0c3166a4c32a5bca6d3a64bff6e0"
	fmt.Println("requestApiKey")
	fmt.Println(requestApiKey)
	fmt.Println("serverApiKey")
	fmt.Println(serverApiKey)


	if requestApiKey != serverApiKey {
		return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "unauthorized",
		})
	}

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
