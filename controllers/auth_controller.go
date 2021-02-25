package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/services"
)



func LoginClient(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	token, err := services.LoginClient(email, password)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(fiber.Map{"token": token})
}

func LoginEmployee(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	token, err := services.LoginClient(email, password)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(fiber.Map{"token": token})
}

