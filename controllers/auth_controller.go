package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/services"
)

func LoginClient(context *fiber.Ctx) error {
	email := context.FormValue("email")
	password := context.FormValue("password")

	token, err := services.LoginClient(email, password)
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.JSON(fiber.Map{"token": token})
}

func LoginEmployee(context *fiber.Ctx) error {
	email := context.FormValue("email")
	password := context.FormValue("password")
	token, err := services.LoginEmployer(email, password)
	if err != nil {
		return context.SendStatus(fiber.StatusNotFound)
	}
	return context.JSON(fiber.Map{"token": token})
}
