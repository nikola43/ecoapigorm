package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func SignUpClient(context *fiber.Ctx) error {
	return context.JSON(fiber.Map{"message": "SignUpClient"})
}

func SignUpEmployee(context *fiber.Ctx) error {
	return context.JSON(fiber.Map{"message": "SignUpEmployee"})
}

