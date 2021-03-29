package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models/promos"
	"github.com/nikola43/ecoapigorm/services"
	"strconv"
)

func CreatePromo(context *fiber.Ctx) error {
	createPromoRequest := new(promos.CreatePromoRequest)
	err := context.BodyParser(&createPromoRequest)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := services.CreatePromoService(createPromoRequest)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(response)

}
func GetPromosController(context *fiber.Ctx) error {

	promos, err := services.GetAllPromos()
	if err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(promos)
}

func DeletePromoByID(context *fiber.Ctx) error {
	promoID, _ := strconv.ParseUint(context.Params("promo_id"), 10, 64)

	err := services.DeletePromoByID(uint(promoID))
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

