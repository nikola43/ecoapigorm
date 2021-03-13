package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/services"
)

func LoginClient(context *fiber.Ctx) error {
	clientLoginRequest := new(models.LoginClientRequest)

	err := context.BodyParser(clientLoginRequest)
	if err != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	clientLoginResponse, err := services.LoginClient(clientLoginRequest.Email, clientLoginRequest.Password)
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": errors.New("not found"),
		})
	}

	return context.JSON(clientLoginResponse)
}

func LoginEmployee(context *fiber.Ctx) error {
	loginEmployeeRequest := new(models.LoginEmployeeRequest)

	err := context.BodyParser(loginEmployeeRequest)
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	clientEmployeeResponse, err := services.LoginEmployee(loginEmployeeRequest.Email, loginEmployeeRequest.Password)
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.JSON(clientEmployeeResponse)
}
