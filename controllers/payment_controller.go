package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models/payments"
	"github.com/nikola43/ecoapigorm/services"
)

func CreatePayment(context *fiber.Ctx) error {
	// todo validate company, clinic y employee

	createPaymentRequest := new(payments.CreatePaymentRequest)

	// parse request
	err := context.BodyParser(createPaymentRequest)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	payment, err := services.CreatePayment(createPaymentRequest)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(payment)

}
func ValidatePayment(context *fiber.Ctx) error {
	sessionID := context.Params("session_id")
	payment, err := services.ValidatePayment(sessionID)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}


	return context.Status(fiber.StatusOK).JSON(payment)
}
