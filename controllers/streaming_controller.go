package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models/streaming"
	streamings "github.com/nikola43/ecoapigorm/models/streamings"
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

func CreateStreaming(context *fiber.Ctx) error {
	createStreamingRequest := new(streamings.CreateStreamingRequest)
	fmt.Println(createStreamingRequest)


	// parse request
	 err := context.BodyParser(createStreamingRequest)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	createStreamingResponse, err := services.CreateStreaming(createStreamingRequest)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println(createStreamingResponse)
	return context.Status(fiber.StatusOK).JSON(createStreamingResponse)
}
