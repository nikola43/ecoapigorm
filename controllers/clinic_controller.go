package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/clients"
	clinicModels "github.com/nikola43/ecoapigorm/models/clinic"
	"github.com/nikola43/ecoapigorm/services"
	"strconv"
)

func GetClinicById(context *fiber.Ctx) error {
	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)
	fmt.Println(clinicID)
	if clinic, err := services.GetClinicByID(uint(clinicID))
		err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "clinic not found",
		})
	} else {
		return context.JSON(clinic)
	}
}

func CreateClinic(context *fiber.Ctx) error {
	createClinicRequest := new(clinicModels.CreateClinicRequest)
	createClinicResponse := new(clinicModels.CreateClinicResponse)
	var err error
	clinic := models.Clinic{}
	employee := models.Employee{}

	// parse request
	if err = context.BodyParser(createClinicRequest);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "bad request",
		})
	}

	// validation ---------------------------------------------------------------------
	v := validator.New()
	err = v.Struct(createClinicRequest)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e != nil {
				return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"error": "validation_error: " + e.Field(),
				})
			}
		}
	}

	// check if employee exist
	GormDBResult := database.GormDB.
		Where("id = ?", createClinicRequest.EmployeeID).
		Find(&employee)

	if GormDBResult.Error != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "internal server",
		})
	}

	if employee.ID < 1 {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "employee id not found",
		})
	}

	// check if clinic already exist
	GormDBResult = database.GormDB.
		Where("name = ? AND employee_id = ?", createClinicRequest.Name, createClinicRequest.EmployeeID).
		Find(&clinic)

	if GormDBResult.Error != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "internal server",
		})
	}

	if clinic.ID > 0 {
		return context.Status(fiber.StatusConflict).JSON(&fiber.Map{
			"error": "employee id already associated to other clinic",
		})
	}

	// create and response
	if createClinicResponse, err = services.CreateClinic(createClinicRequest);
		err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	} else {
		return context.JSON(createClinicResponse)
	}
}

func GetClientsByClinicID(context *fiber.Ctx) error {
	clientsList := make([]clients.ListClientRequest, 0)
	var err error

	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)
	fmt.Println(clinicID)
	if clientsList, err = services.GetClientsByClinicID(uint(clinicID))
		err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "clinic not found",
		})
	} else {
		return context.JSON(clientsList)
	}
}
