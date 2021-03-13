package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	modelsEmployees "github.com/nikola43/ecoapigorm/models/employee"
	"github.com/nikola43/ecoapigorm/services"
	"strconv"
)

func CreateEmployee(context *fiber.Ctx) error {
	var err error

	createEmployeeRequest := new(modelsEmployees.CreateEmployeeRequest)
	createEmployeeResponse := new(modelsEmployees.CreateEmployeeResponse)

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

func GetEmployeesByParentEmployeeID(context *fiber.Ctx) error {
	parentEmployeeID, _ := strconv.ParseUint(context.Params("parent_employee_id"), 10, 64)

	employees := make([]models.Employee, 0)
	var err error

	if employees, err = services.GetEmployeesByParentEmployeeID(uint(parentEmployeeID)); err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(employees)
}

func BuyCredits(context *fiber.Ctx) error {
	sessionID := context.Params("session_id")
	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)


	payment, err := services.BuyCredits(sessionID, uint(clinicID))
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"credits": payment.Amount,
	})
}

func Invite(context *fiber.Ctx) error {
	var err error
	var employees = make([]models.Employee, 0)

	// parse request
	if err = context.BodyParser(&employees);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	services.Invite(employees)

	return context.SendStatus(fiber.StatusOK)
}
