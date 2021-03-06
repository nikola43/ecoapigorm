package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/services"
)

func LoginClient(context *fiber.Ctx) error {
	clientLoginRequest := new(models.LoginClientRequest)
	clientLoginResponse := new(models.LoginClientResponse)
	var err error

	if err = context.BodyParser(clientLoginRequest); err != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	if clientLoginResponse, err = services.LoginClient(clientLoginRequest.Email, clientLoginRequest.Password); err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": errors.New("not found"),
		})
	}

	return context.JSON(clientLoginResponse)
}

func LoginEmployee(context *fiber.Ctx) error {
	clientEmployeeRequest := new(models.LoginEmployeeRequest)
	clientEmployeeResponse := new(models.LoginEmployeeResponse)
	var err error

	if err = context.BodyParser(clientEmployeeRequest); err != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	if clientEmployeeResponse, err = services.LoginEmployee(clientEmployeeRequest.Email, clientEmployeeRequest.Password); err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": errors.New("not found"),
		})
	}

	return context.JSON(clientEmployeeResponse)
}
