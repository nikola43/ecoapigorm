package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	recovery "github.com/nikola43/ecoapigorm/models/recovery"
	"github.com/nikola43/ecoapigorm/services"
	"github.com/nikola43/ecoapigorm/utils"
)

func PassRecoveryClient(context *fiber.Ctx) error {
	passRecoveryClientRequest := new(recovery.PassRecoveryRequest)
	var err error

	if err = context.BodyParser(passRecoveryClientRequest);
		err != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	err = services.PassRecoveryClientService(passRecoveryClientRequest)

	if err != nil {
		return context.SendStatus(fiber.StatusNotFound)
	}


	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func PassRecoveryEmployee(context *fiber.Ctx) error {
	passRecoveryClientRequest := new(recovery.PassRecoveryRequest)
	var err error

	if err = context.BodyParser(passRecoveryClientRequest);
		err != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	err = services.PassRecoveryEmployeeService(passRecoveryClientRequest)

	if err != nil {
		return context.SendStatus(fiber.StatusNotFound)
	}


	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func ValidateRecovery(context *fiber.Ctx) error {
	recoveryToken := context.Params("recovery_token")
	fmt.Println("recoveryToken")
	fmt.Println(recoveryToken)

	userRecovery := new(recovery.UserRecovery)
	utils.GetModelByField(userRecovery, "token", recoveryToken)

	// todo check claims
	if userRecovery.ID < 1 {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "invalid recovery token",
		})
	}

	return context.Status(fiber.StatusOK).JSON(userRecovery)
}
