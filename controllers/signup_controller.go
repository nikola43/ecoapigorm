package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/services"
)



func SignUpClient(context *fiber.Ctx) error {
	name := context.FormValue("name")
	email := context.FormValue("email")
	password := context.FormValue("password")

	token, err := services.SignUpClient(name, email, password)
	if err != nil {
		return context.SendStatus(fiber.StatusNotFound)
	}

	return context.JSON(fiber.Map{"token": token})
}

func SignUpEmployee(context *fiber.Ctx) error {
	email := context.FormValue("email")
	password := context.FormValue("password")
	token, err := services.LoginClient(email, password)
	if err != nil {
		return context.SendStatus(fiber.StatusNotFound)
	}
	return context.JSON(fiber.Map{"token": token})
}

