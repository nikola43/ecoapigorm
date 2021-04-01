package controllers

import (
	"github.com/gofiber/fiber/v2"
	recovery "github.com/nikola43/ecoapigorm/models/recovery"
	"github.com/nikola43/ecoapigorm/services"
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

	return context.SendStatus(fiber.StatusOK)
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

	return context.SendStatus(fiber.StatusOK)
}
