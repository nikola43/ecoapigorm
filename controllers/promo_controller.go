package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/services"
)

func CreatePromo(context *fiber.Ctx) error {
	return context.SendStatus(fiber.StatusBadRequest)

}
func GetPromosController(context *fiber.Ctx) error {

	promos,err := services.GetAllPromos()
	if err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(promos)
}
