package controllers

import (
	streamingModels "github.com/nikola43/ecoapigorm/models/streamings"
	"github.com/nikola43/ecoapigorm/utils"
	//"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/nikola43/ecoapigorm/models"
	modelsClient "github.com/nikola43/ecoapigorm/models/clients"
	"github.com/nikola43/ecoapigorm/services"
	"strconv"
	//"strings"
)

func NotifyClient(context *fiber.Ctx) error {
	client := new(models.Client)
	var err error

	// parse request
	if err = context.BodyParser(client);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	sendEmailManager := utils.SendEmailManager{
		ToEmail:         client.Email,
	}

	sendEmailManager.SendMail("notify.html", "Nuevo contenido disponible")

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func GetClientClinicIDByEmail(context *fiber.Ctx) error {
	clientEmail := context.Params("client_email")
	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)

	client, err := services.GetClientClinicIDByEmail(uint(clinicID), clientEmail)
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

func GetAllImagesByClientAndClinicID(context *fiber.Ctx) error {
	var err error
	images := make([]models.Image, 0)
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)
	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)

	if images, err = services.GetAllImagesByClientAndClinicID(uint(clientID), uint(clinicID)); err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(images)
}

func GetAllStreamingByClientID(context *fiber.Ctx) error {
	clientID := context.Params("client_id")
	videos := make([]streamingModels.Streaming, 0)
	var err error

	if videos, err = services.GetAllStreamingByClientID(clientID); err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(videos)
}

func GetAllStreamingByClientANDClinicID(context *fiber.Ctx) error {
	clientID := context.Params("client_id")
	clinicID := context.Params("clinic_id")
	videos := make([]streamingModels.Streaming, 0)
	var err error

	if videos, err = services.GetAllStreamingByClientANDClinicID(clientID, clinicID); err != nil {
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

func GetAllVideosByClientAndClinicID(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)
	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)
	videos := make([]models.Video, 0)
	var err error

	if videos, err = services.GetAllVideosByClientAndClinicID(uint(clientID), uint(clinicID)); err != nil {
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

func GetHeartbeatByClientAndClinicID(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)
	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)

	heartbeat := &models.Heartbeat{}
	var err error

	if heartbeat, err = services.GetHeartbeatByClientAndClinicID(uint(clientID), uint(clinicID)); err != nil {
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
	if createClientResponse, err = services.CreateClientFromApp(createClientFromAppRequest);
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

func IncrementDiskQuoteLevel(context *fiber.Ctx) error {
	listClientResponse := new(modelsClient.ListClientResponse)

	if err := context.BodyParser(listClientResponse);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	employeeTokenClaims, err := utils.GetEmployeeTokenClaims(context)
	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	err = services.IncrementDiskQuoteLevel(employeeTokenClaims.ClinicID, listClientResponse.ID)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func ChangePassClient(context *fiber.Ctx) error {
	changePassClientRequest := new(modelsClient.ChangePassClientRequest)
	var err error

	if err = context.BodyParser(changePassClientRequest);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	err = services.ChangePassClientService(changePassClientRequest)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func UpdateClient(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)
	updateClientRequest := new(modelsClient.UpdateClientRequest)

	err := context.BodyParser(updateClientRequest)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	client, err := services.UpdateClientService(uint(clientID), updateClientRequest)
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(client)
}

func UnassignClientByID(context *fiber.Ctx) error {
	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)

	// TODO Cualquier propietario podría borrar usuarios de otras clínicas
	/*
		employeeTokenClaims, err := utils.GetEmployeeTokenClaims(context)
		if err != nil {
			return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}
	*/

	err := services.UnassignClientByID(uint(clinicID), uint(clientID))
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func RefreshClient(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)

	client, err := services.RefreshClient(uint(clientID))
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(client)
}
