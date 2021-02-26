package middleware

import (
	"github.com/gofiber/fiber/v2"
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


