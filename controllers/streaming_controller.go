package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	streamings "github.com/nikola43/ecoapigorm/models/streamings"
	"github.com/nikola43/ecoapigorm/services"
	"github.com/nikola43/ecoapigorm/utils"
	"strconv"
)

func GetStreamingByCodeController(context *fiber.Ctx) error {
	code := context.Params("code")
	streaming, err := services.GetStreamingByCodeService(code)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusNotFound, err, context)
	}

	return context.Status(fiber.StatusOK).JSON(streaming)
}

func CreateStreaming(context *fiber.Ctx) error {
	employeeTokenClaims, getEmployeeTokenClaimsErr := utils.GetEmployeeTokenClaims(context)
	if getEmployeeTokenClaimsErr != nil {
		return utils.ReturnErrorResponse(fiber.StatusUnauthorized, getEmployeeTokenClaimsErr, context)
	}
	fmt.Println(employeeTokenClaims)


	createStreamingRequest := new(streamings.CreateStreamingRequest)

	err := utils.ParseAndValidate(context, createStreamingRequest)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	createStreamingResponse, createErr := services.CreateStreaming(createStreamingRequest)
	if createErr != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, createErr, context)
	}

	return context.Status(fiber.StatusOK).JSON(createStreamingResponse)
}

func DeleteStreamingByID(context *fiber.Ctx) error {
	streamingID, _ := strconv.ParseUint(context.Params("streaming_id"), 10, 64)

	err := services.DeleteStreamingByID(uint(streamingID))
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	return utils.ReturnSuccessResponse(context)
}

func UpdateStreaming(context *fiber.Ctx) error {
	streaming := new(streamings.Streaming)

	err := utils.ParseAndValidate(context, streaming)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	streaming, err = services.UpdateStreaming(streaming)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	return context.Status(fiber.StatusOK).JSON(streaming)
}
