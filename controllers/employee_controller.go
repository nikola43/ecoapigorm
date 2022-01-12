package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	modelsEmployees "github.com/nikola43/ecoapigorm/models/employee"
	"github.com/nikola43/ecoapigorm/services"
	"github.com/nikola43/ecoapigorm/utils"
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
			"error": err.Error(),
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

	// check if token exist
	invitation := new(models.Invitation)
	result := database.GormDB.Where("to_email = ? AND token = ?", createEmployeeRequest.Email, createEmployeeRequest.Token).Find(&invitation)
	if result.Error != nil {
		return result.Error
	}

	// create and response
	if createEmployeeResponse, err = services.CreateEmployeeFromPanel(createEmployeeRequest); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "internal server",
		})
	}

	database.GormDB.Delete(invitation)

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

	var employees = make([]models.Employee, 0)
	employeeTokenClaims, err := utils.GetEmployeeTokenClaims(context)
	fmt.Println("Invite")
	fmt.Println("employeeTokenClaims")
	fmt.Println(employeeTokenClaims)
	fmt.Println(employeeTokenClaims.CompanyID)

	// parse request
	err = context.BodyParser(&employees)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	err = services.Invite(employeeTokenClaims, employees)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func GetCompaniesByEmployeeID(context *fiber.Ctx) error {
	employeeID, _ := strconv.ParseUint(context.Params("employee_id"), 10, 64)

	var err error

	list, err := services.GetCompaniesByEmployeeID(uint(employeeID))

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(list)
}

func DeleteEmployeeByEmployeeID(context *fiber.Ctx) error {
	employeeID, _ := strconv.ParseUint(context.Params("employee_id"), 10, 64)
	employeeTokenClaims, err := utils.GetEmployeeTokenClaims(context)
	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	err = services.DeleteEmployeeByEmployeeID(employeeTokenClaims.ID, uint(employeeID))
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func ValidateInvitation(context *fiber.Ctx) error {
	invitationToken := context.Params("invitation_token")
	fmt.Println(invitationToken)

	invitation := new(models.Invitation)
	utils.GetModelByField(invitation, "token", invitationToken)

	if invitation.ID < 1 {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "invalid invitation token",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success":            true,
		"clinic_id":          invitation.FromClinicID,
		"company_id":         invitation.CompanyID,
		"employee_email":     invitation.ToEmail,
		"token":              invitationToken,
		"parent_employee_id": invitation.ParentEmployeeID,
	})
}

func ChangePassEmployee(context *fiber.Ctx) error {
	changePasswordEmployeeRequest := new(modelsEmployees.ChangePasswordEmployeeRequest)
	var err error

	if err = context.BodyParser(changePasswordEmployeeRequest);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	err = services.ChangePassEmployeeService(changePasswordEmployeeRequest)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}
