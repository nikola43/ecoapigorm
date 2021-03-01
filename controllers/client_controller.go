package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models"
	modelsClient "github.com/nikola43/ecoapigorm/models/clients"
	"github.com/nikola43/ecoapigorm/services"
)

func GetAllImagesByClientID(context *fiber.Ctx) error {
	clientID := context.Params("client_id")
	images := make([]models.Image, 0)
	var err error

	if images, err = services.GetAllImagesByClientID(clientID); err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(images)
}

func CreateClient(context *fiber.Ctx) error {
	createClientRequest := new(modelsClient.CreateClientRequest)
	createClientResponse := new(modelsClient.CreateClientResponse)
	var err error

	if err = context.BodyParser(createClientRequest);
	err != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	if createClientResponse, err = services.CreateClient(createClientRequest);
	err != nil {
	return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
		"error": errors.New("not found"),
	})
	}else {
		return context.JSON(createClientResponse)
	}
}

func ChangePassClient(context *fiber.Ctx) error {
	changePassClientRequest := new(modelsClient.ChangePassClientRequest)
	var err error

	if err = context.BodyParser(changePassClientRequest);
	err != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	err = services.ChangePassClientService(changePassClientRequest)

	if err != nil {
		return context.SendStatus(fiber.StatusNotFound)
	}

	return context.SendStatus(fiber.StatusOK)
}

