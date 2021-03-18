package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/awsmanager"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	companyModels "github.com/nikola43/ecoapigorm/models/company"
	"github.com/nikola43/ecoapigorm/services"
	"github.com/nikola43/ecoapigorm/utils"
	"strconv"
	"strings"
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
	company := models.Company{}
	employee := models.Employee{}

	employeeTokenClaims, err := utils.GetEmployeeTokenClaims(context)
	if err != nil {
		return context.SendStatus(fiber.StatusUnauthorized)
	}

	fmt.Println("employeeTokenClaims")
	fmt.Println(employeeTokenClaims)

	// parse request
	if err = context.BodyParser(createCompanyRequest);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
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
		Where("id = ?", employeeTokenClaims.ID).
		Find(&employee)

	if GormDBResult.Error != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "internal server",
		})
	}

	if employeeTokenClaims.ID < 1 {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "employee id not found",
		})
	}

	// check if Company already exist
	GormDBResult = database.GormDB.
		Where("name = ?", createCompanyRequest.Name).
		Find(&company)

	if GormDBResult.Error != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": GormDBResult.Error.Error(),
		})
	}

	if company.ID > 0 {
		return context.Status(fiber.StatusConflict).JSON(&fiber.Map{
			"error": "company already exist",
		})
	}

	// create and response
	if createCompanyResponse, err = services.CreateCompany(employeeTokenClaims.ID, createCompanyRequest);
		err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	employeeToken, err := utils.GenerateEmployeeToken(
		employee.Name,
		createCompanyResponse.ID,
		employee.Clinic.ID,
		employee.ID,
		employee.Email,
		createCompanyResponse.Name,
		employee.Clinic.Name,
		employee.Role)

	createCompanyResponse.Token = employeeToken

	// create bucket for company
	bucketName := strings.ToLower(strings.ReplaceAll(createCompanyRequest.Name, " ", "_"))
	err = awsmanager.AwsManager.CreateBucket(strings.ToLower(bucketName))
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.JSON(createCompanyResponse)

}

func GetEmployeesByCompanyID(context *fiber.Ctx) error {
	companyID, _ := strconv.ParseUint(context.Params("company_id"), 10, 64)

	if list, err := services.GetEmployeesByCompanyID(uint(companyID)); err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "Company not found",
		})
	} else {
		return context.JSON(list)
	}
}

func GetClinicsByCompanyID(context *fiber.Ctx) error {
	companyID, _ := strconv.ParseUint(context.Params("company_id"), 10, 64)

	// todo validad solo puede verlo el dueño
	if list, err := services.GetClinicsByCompanyID(uint(companyID)); err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	} else {
		return context.JSON(list)
	}
}
