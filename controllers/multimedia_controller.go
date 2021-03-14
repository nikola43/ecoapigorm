package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/services"
	"strconv"
)

func DeleteImage(context *fiber.Ctx) error {
	id, _ := strconv.ParseUint(context.Params("id"), 10, 64)

	err := services.DeleteImage(uint(id))
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println(id)

	return context.SendStatus(fiber.StatusOK)

}

func DeleteVideo(context *fiber.Ctx) error {
	id, _ := strconv.ParseUint(context.Params("video_id"), 10, 64)

	fmt.Println(id)

	return context.SendStatus(fiber.StatusBadRequest)

}
func DeleteHolographic(context *fiber.Ctx) error {
	id, _ := strconv.ParseUint(context.Params("holopraphic_id"), 10, 64)

	fmt.Println(id)

	return context.SendStatus(fiber.StatusBadRequest)

}

func DeleteHeartbeat(context *fiber.Ctx) error {
	id, _ := strconv.ParseUint(context.Params("heartbeat_id"), 10, 64)

	fmt.Println(id)

	return context.SendStatus(fiber.StatusBadRequest)

}
