package utils

import "github.com/gofiber/fiber/v2"

func ReturnErrorResponse(status int, err error, context *fiber.Ctx) error {
	return context.Status(status).JSON(&fiber.Map{
		"error": err.Error(),
	})
}

func ReturnSuccessResponse(context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}