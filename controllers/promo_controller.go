package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	companyModels "github.com/nikola43/ecoapigorm/models/company"
	"github.com/nikola43/ecoapigorm/services"
)

func CreatePromo(context *fiber.Ctx) error {
	createPromoRequest := new(companyModels.CreateCompanyRequest)
	createPromoResponse := new(companyModels.CreateCompanyResponse)
	var err error
	Company := models.Company{}
	employee := models.Employee{}

	// parse request
	if err = context.BodyParser(createPromoRequest);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "bad request",
		})
	}

	// validation ---------------------------------------------------------------------
	v := validator.New()
	err = v.Struct(createPromoRequest)
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
		Where("id = ?", createPromoRequest.EmployeeID).
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

	// check if Company already exist
	GormDBResult = database.GormDB.
		Where("name = ? AND employee_id = ?", createPromoRequest.Name, createPromoRequest.EmployeeID).
		Find(&Company)

	if GormDBResult.Error != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "internal server",
		})
	}

	if Company.ID > 0 {
		return context.Status(fiber.StatusConflict).JSON(&fiber.Map{
			"error": "employee id already associated to other Company",
		})
	}

	// create and response
	if createPromoResponse, err = services.CreateCompany(createPromoRequest);
		err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	} else {
		return context.JSON(createPromoResponse)
	}
}
func GetPromosController(context *fiber.Ctx) error {

	promos,err := services.GetAllPromos()
	if err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(promos)
}
