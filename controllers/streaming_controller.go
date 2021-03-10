package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models/streaming"
	"github.com/nikola43/ecoapigorm/services"
)

func GetStreamingByCodeController(context *fiber.Ctx) error {
	code := context.Params("code")
	var streaming = streaming.Streaming{}
	var err error

	if streaming, err = services.GetStreamingByCodeService(code);
	err != nil {
		return context.SendStatus(fiber.StatusNotFound)
	}

	return context.Status(fiber.StatusOK).JSON(streaming)
}
