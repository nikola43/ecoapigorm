package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models/kicks"
	"github.com/nikola43/ecoapigorm/services"
	"strconv"
)

func AddKickToClient(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)
	addKickRequest := new(kicks.AddKickRequest)

	addKickRequest.ClientId = uint(clientID)
	var err error

	// parse request
	if err = context.BodyParser(addKickRequest);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "bad request",
		})
	}

	kick,err := services.CreateNewKickService(*addKickRequest)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(err)
	}

	return context.Status(fiber.StatusCreated).JSON(kick)
}

func GetKicksByClient(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)

	kicks,err := services.GetAllKicksByClientService(uint(clientID))
	if err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(kicks)
}

func DeleteKick(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)
	kickID, _ := strconv.ParseUint(context.Params("kick_id"), 10, 64)

	err := services.DeleteKickByIdService(uint(clientID), uint(kickID))
	if err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.SendStatus(fiber.StatusOK)
}

func ResetAllByClientKicks(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)

	err := services.ResetAllKicksByClientService(uint(clientID))
	if err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.SendStatus(fiber.StatusOK)
}
