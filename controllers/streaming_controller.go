package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/services"
)

func GetStreamingByCode(context *fiber.Ctx) error {
	code := context.Params("code")
	fmt.Println(code)
	var streaming = &models.Streaming{}
	var err error

	streaming, err = services.GetStreamingByCode(code)
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(streaming)
}
