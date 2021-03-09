package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	modelsClient "github.com/nikola43/ecoapigorm/models/clients"
	"github.com/nikola43/ecoapigorm/models/streaming"
	"github.com/nikola43/ecoapigorm/services"
	"strconv"
)

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

	// parse request
	if err = context.BodyParser(createClientRequest);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "bad request",
		})
	}

	// validation ---------------------------------------------------------------------
	v := validator.New()
	err = v.Struct(createClientRequest)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e != nil {
				return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"error": "validation_error: " + e.Field(),
				})
			}
		}
	}

	// check if client already exist
	client := models.Client{}
	GormDBResult := database.GormDB.
		Where("email = ?", createClientRequest.Email).
		Find(&client)

	if GormDBResult.Error != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "internal server",
		})
	}

	// create and response
	if createClientResponse, err = services.CreateClient(createClientRequest);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	} else {
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

func UploadMultimedia(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)
	uploadMode, _ := strconv.ParseUint(context.Params("upload_mode"), 10, 64)
	//uploadMode, _ := strconv.ParseUint(context.FormValue("upload_mode"), 10, 64)
	fmt.Println(uploadMode)

	// Get first file from form field "document":
	uploadedFile, err := context.FormFile("file")
	//uploadedFileMode, err := context.FormFile("mode")
	if err != nil {
		return err
	}


	err = services.UploadMultimedia(context,uint(clientID), uploadedFile)
	if err != nil {
		return err
	}


	return context.SendStatus(fiber.StatusOK)
}
