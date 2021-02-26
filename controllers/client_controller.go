package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models"
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
