package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllClientsByClinicId(context *fiber.Ctx) error {
	return context.SendString("Hello, World!")
}
