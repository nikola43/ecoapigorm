package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/services"
	"github.com/nikola43/ecoapigorm/utils"
)


func LoginClient(context *fiber.Ctx) error {
	clientLoginRequest := new(models.LoginClientRequest)

	err := context.BodyParser(clientLoginRequest)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	clientLoginResponse, err := services.LoginClient(clientLoginRequest.Email, clientLoginRequest.Password)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusNotFound, err, context)
	}

	return context.JSON(clientLoginResponse)
}

func LoginEmployee(context *fiber.Ctx) error {
	loginEmployeeRequest := new(models.LoginEmployeeRequest)

	err := context.BodyParser(loginEmployeeRequest)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusNotFound, err, context)
	}

	clientEmployeeResponse, err := services.LoginEmployee(loginEmployeeRequest.Email, loginEmployeeRequest.Password)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusNotFound, err, context)
	}

	return context.JSON(clientEmployeeResponse)
}
