package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/services"
	"strconv"
)

func GetCalculatorByWeek(context *fiber.Ctx) error {
	var calculatorDetail = &models.CalculatorDetail{}
	var err error
	week, _ := strconv.ParseUint(context.Params("week"), 10, 64)

	if calculatorDetail, err = services.GetCalculatorByWeekNumber(uint(week)); err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": errors.New("not found"),
		})
	}
	return context.JSON(calculatorDetail)

}
