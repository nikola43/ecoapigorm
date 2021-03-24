package controllers

import (
	"github.com/nikola43/ecoapigorm/utils"
	//"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/nikola43/ecoapigorm/models"
	modelsClient "github.com/nikola43/ecoapigorm/models/clients"
	"github.com/nikola43/ecoapigorm/models/streaming"
	"github.com/nikola43/ecoapigorm/services"
	"strconv"
	//"strings"
)

func GetClientById(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)

	client, err := services.GetClientById(uint(clientID))
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(client)

}

func GetClientByEmail(context *fiber.Ctx) error {
	clientEmail := context.Params("client_email")

	client, err := services.GetClientByEmail(clientEmail)
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return context.Status(fiber.StatusOK).JSON(client)
}

func GetAllImagesByClientID(context *fiber.Ctx) error {
	var err error
	images := make([]models.Image, 0)
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)

	if images, err = services.GetAllImagesByClientID(uint(clientID)); err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(images)
}

func GetAllStreamingByClientID(context *fiber.Ctx) error {
	clientID := context.Params("client_id")
	videos := make([]streaming.Streaming, 0)
	var err error

	if videos, err = services.GetAllStreamingByClientID(clientID); err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(videos)
}

func GetAllVideosByClientID(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)
	videos := make([]models.Video, 0)
	var err error

	if videos, err = services.GetAllVideosByClientID(uint(clientID)); err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(videos)
}

func GetAllHolographicsByClientID(context *fiber.Ctx) error {
	clientID := context.Params("client_id")
	holographics := make([]models.Holographic, 0)
	var err error

	if holographics, err = services.GetAllHolographicsByClientID(clientID); err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(holographics)
}

func GetHeartbeatByClientID(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)

	heartbeat := &models.Heartbeat{}
	var err error

	if heartbeat, err = services.GetHeartbeatByClientID(uint(clientID)); err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(heartbeat)
}

func CreateClient(context *fiber.Ctx) error {
	createClientFromAppRequest := new(modelsClient.CreateClientFromAppRequest)
	createClientResponse := new(modelsClient.CreateClientResponse)
	var err error

	// parse request
	if err = context.BodyParser(createClientFromAppRequest);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	// validation ---------------------------------------------------------------------
	v := validator.New()
	err = v.Struct(createClientFromAppRequest)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e != nil {
				return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"error": "validation_error: " + e.Field(),
				})
			}
		}
	}

	// create and response
	if createClientResponse, err = services.CreateClient(createClientFromAppRequest);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	} else {
		welcomeEmail := utils.SendEmailManager{ToEmail: createClientFromAppRequest.Email,
			ToName: createClientFromAppRequest.Name,
		}
		welcomeEmail.SendMail("welcome.html", "Bienvenido "+createClientFromAppRequest.Name)
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

func UpdateClient(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)
	updateClientRequest := new(modelsClient.UpdateClientRequest)
	var err error

	if err = context.BodyParser(updateClientRequest);
		err != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	err = services.UpdateClientService(uint(clientID), updateClientRequest)

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

func PasswordRecovery(context *fiber.Ctx) error {
	return context.SendStatus(fiber.StatusOK)
}

// TODO ¿este metodo es necesario?
func DeleteClientByID(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)

	err := services.DeleteClientByID(uint(clientID))
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}
