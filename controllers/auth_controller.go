package controllers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/services"
)

func LoginClient(context *fiber.Ctx) error {
	clientLoginRequest := new(models.ClientLoginRequest)
	clientLoginResponse := new(models.ClientLoginResponse)
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
	fmt.Println("LoginEmployee")
	/*
		email := context.FormValue("email")
		password := context.FormValue("password")
		token, err := services.LoginEmployer(email, password)
		if err != nil {
			return context.SendStatus(fiber.StatusNotFound)
		}
	*/
	return context.JSON(fiber.Map{"token": ""})
}
