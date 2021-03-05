package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	companyModels "github.com/nikola43/ecoapigorm/models/company"
	"github.com/nikola43/ecoapigorm/services"
	"strconv"
)

func GetCompanyById(context *fiber.Ctx) error {
	var err error
	company := &models.Company{}
	companyID, _ := strconv.ParseUint(context.Params("company_id"), 10, 64)

	if company, err = services.GetCompanyByID(uint(companyID)); err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "Company not found",
		})
	}

	return context.JSON(company)
}

func CreateCompany(context *fiber.Ctx) error {
	createCompanyRequest := new(companyModels.CreateCompanyRequest)
	createCompanyResponse := new(companyModels.CreateCompanyResponse)
	var err error
	Company := models.Company{}
	employee := models.Employee{}

	// parse request
	if err = context.BodyParser(createCompanyRequest);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "bad request",
		})
	}

	// validation ---------------------------------------------------------------------
	v := validator.New()
	err = v.Struct(createCompanyRequest)
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
		Where("id = ?", createCompanyRequest.EmployeeID).
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
		Where("name = ? AND employee_id = ?", createCompanyRequest.Name, createCompanyRequest.EmployeeID).
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
	if createCompanyResponse, err = services.CreateCompany(createCompanyRequest);
		err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	} else {
		return context.JSON(createCompanyResponse)
	}
}

func GetEmployeesByCompanyID(context *fiber.Ctx) error {

	list := make([]models.Employee, 0)
	var err error

	companyID, _ := strconv.ParseUint(context.Params("company_id"), 10, 64)
	fmt.Println(companyID)
	if list, err = services.GetEmployeesByCompanyID(uint(companyID))
		err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "Company not found",
		})
	}

	return context.JSON(list)
}
