package controllers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models"
	modelsClient "github.com/nikola43/ecoapigorm/models/clients"
	"github.com/nikola43/ecoapigorm/services"
	"github.com/nikola43/ecoapigorm/utils"
)

func GetAllImagesByClientID(context *fiber.Ctx) error {
	var err error
	var tokenClaims = models.ClientTokenClaims{}
	images := make([]models.Image, 0)
	clientID := context.Params("client_id")

	// todo example get token claims
	tokenClaims, err = utils.GetTokenClaims(context)
	fmt.Println(tokenClaims)

	if images, err = services.GetAllImagesByClientID(clientID); err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(images)
}

func GetAllVideosByClientID(context *fiber.Ctx) error {
	clientID := context.Params("client_id")
	videos := make([]models.Video, 0)
	var err error

	if videos, err = services.GetAllVideosByClientID(clientID); err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(videos)
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

func PassRecoveryClient(context *fiber.Ctx) error {
	passRecoveryClientRequest := new(modelsClient.PassRecoveryRequest)
	var err error

	if err = context.BodyParser(passRecoveryClientRequest);
	err != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	err = services.PassRecoveryClientService(passRecoveryClientRequest)

	if err != nil {
		return context.SendStatus(fiber.StatusNotFound)
	}

	return context.SendStatus(fiber.StatusOK)
}

