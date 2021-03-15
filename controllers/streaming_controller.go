package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models/streaming"
	streamings "github.com/nikola43/ecoapigorm/models/streamings"
	"github.com/nikola43/ecoapigorm/services"
	"strconv"
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



func DeleteStreamingByID(context *fiber.Ctx) error {
	streamingID, _ := strconv.ParseUint(context.Params("streaming_id"), 10, 64)

	err := services.DeleteStreamingByID(uint(streamingID))
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}
