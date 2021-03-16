package services

import (
	"errors"
	"fmt"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	modelsEmployees "github.com/nikola43/ecoapigorm/models/employee"
	"github.com/nikola43/ecoapigorm/utils"
)

func CreateEmployeeFromPanel(createEmployeeRequest *modelsEmployees.CreateEmployeeRequest) (*modelsEmployees.CreateEmployeeResponse, error) {

	// todo if has clinic id, check if clinic exist

	// todo cambiar role por constantes
	// create newEmployee on DB
	newEmployee := models.Employee{
		ParentEmployeeID: createEmployeeRequest.ParentEmployeeID,
		ClinicID:         createEmployeeRequest.ClinicID,
		Name:             createEmployeeRequest.Name,
		Email:            createEmployeeRequest.Email,
		LastName:         createEmployeeRequest.LastName,
		Role:             "employee",
		Password:         utils.HashPassword([]byte(createEmployeeRequest.Password)),
	}

	fmt.Println("createEmployeeRequest")
	fmt.Println(createEmployeeRequest)

	result := database.GormDB.Create(&newEmployee)
	if result.Error != nil {
		return nil, result.Error
	}

	// generate response
	createEmployeeResponse := modelsEmployees.CreateEmployeeResponse{
		ID:       newEmployee.ID,
		ClinicID: newEmployee.ClinicID,
		Email:    newEmployee.Email,
		Name:     newEmployee.Name,
		LastName: newEmployee.LastName,
		Role:     newEmployee.Role,
	}

	return &createEmployeeResponse, result.Error
}

func GetEmployeesByParentEmployeeID(parentEmployeeID uint) ([]models.Employee, error) {
	var list = make([]models.Employee, 0)

	if err := database.GormDB.Where("parent_employee_id = ?", parentEmployeeID).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func BuyCredits(sessionID string, clinicID uint) (*models.Payment, error) {
	payment := &models.Payment{}

	err := database.GormDB.
		Where("session_id = ? AND clinic_id", sessionID, clinicID).
		Find(&payment).Error

	if err != nil {
		return nil, err
	}

	if payment.ID > 0 {
		return payment, nil
	}

	return nil, errors.New("payment not found")
}

func Invite(employeeTokenClaims *models.EmployeeTokenClaims, employees []models.Employee) error {
	// validation ---------------------------------------------------------------------
	for _, employee := range employees {
		temp := new(models.Employee)
		utils.GetModelByField(temp, "email", employee.Email)
		invitationToken, err := utils.GenerateInvitationToken()
		if err != nil {
			return err
		}
		fmt.Println("employeeTokenClaims.ClinicID")
		fmt.Println(employeeTokenClaims.ClinicID)

		invitation := &models.Invitation{
			ParentEmployeeID: employeeTokenClaims.ID,
			Token:            invitationToken,
			FromEmail:        employeeTokenClaims.Name,
			ToEmail:          employee.Email,
			FromClinicID:     employeeTokenClaims.ClinicID,
		}

		sendEmailManager := utils.SendEmailManager{
			ToEmail:         employee.Email,
			FromName:        employeeTokenClaims.Name,
			ClinicName:      employeeTokenClaims.ClinicName,
			InvitationToken: invitationToken,
		}

		database.GormDB.Create(invitation)

		if temp.ID > 0 {
			sendEmailManager.SendMail("invite_to_clinic.html", employeeTokenClaims.Name+" te ha invitado a su clínica")
		} else {
			sendEmailManager.SendMail("invite_to_register.html", employeeTokenClaims.Name+" te ha invitado a registrarte")
		}
	}

	return nil
}
func DeleteEmployeeByEmployeeID(parentEmployeeID, deletedEmployeeID uint) error {
	deleteEmployee := new(models.Employee)
	deleteEmployeeClinic := new(models.Clinic)

	// check if employee exist
	utils.GetModelByField(deleteEmployee, "id", deletedEmployeeID)
	if deleteEmployee.ID < 1 {
		return errors.New("employee not found")
	}

	// check if employee is owner of any clinic
	utils.GetModelByField(deleteEmployeeClinic, "employee_id", deletedEmployeeID)
	if deleteEmployeeClinic.ID > 1 {
		// update clinic employee id with parent employee id
		database.GormDB.Model(&deleteEmployeeClinic).Update("employee_id", parentEmployeeID)
	}

	// check if employee who make action has deleted employee parent
	if deleteEmployee.ParentEmployeeID != parentEmployeeID {
		// update clinic employee id with parent employee id
		return errors.New("only parent employee can delete employee")
	}

	// delete employee
	result := database.GormDB.Delete(deleteEmployee)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ValidateInvitation(parentEmployeeID, deletedEmployeeID uint) error {

	return nil
}
