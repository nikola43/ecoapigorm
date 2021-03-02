package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	modelsEmployees "github.com/nikola43/ecoapigorm/models/employee"
	"github.com/nikola43/ecoapigorm/services"
)

func CreateEmployee(context *fiber.Ctx) error {
	createEmployeeRequest := new(modelsEmployees.CreateEmployeeRequest)
	createEmployeeResponse := new(modelsEmployees.CreateEmployeeResponse)
	var err error

	// parse request
	if err = context.BodyParser(createEmployeeRequest);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "bad request",
		})
	}

	// validation ---------------------------------------------------------------------
	v := validator.New()
	err = v.Struct(createEmployeeRequest)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e != nil {
				return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"error": "validation_error: " + e.Field(),
				})
			}
		}
	}

	// check if employee already exist
	employee := models.Employee{}
	GormDBResult := database.GormDB.
		Where("email = ?", createEmployeeRequest.Email).
		Find(&employee)

	if GormDBResult.Error != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "internal server",
		})
	}

	if employee.ID > 0 {
		return context.Status(fiber.StatusConflict).JSON(&fiber.Map{
			"error": "employee already exist",
		})
	}

	// create and response
	if createEmployeeResponse, err = services.CreateEmployee(createEmployeeRequest); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "internal server",
		})
	}
	return context.JSON(createEmployeeResponse)
}
