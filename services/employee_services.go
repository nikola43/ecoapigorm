package services

import (
	"errors"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	modelsEmployees "github.com/nikola43/ecoapigorm/models/employee"
	"github.com/nikola43/ecoapigorm/utils"
)

func CreateEmployee(createEmployeeRequest *modelsEmployees.CreateEmployeeRequest) (*modelsEmployees.CreateEmployeeResponse, error) {
	//TODO validate

	employee := models.Employee{
		Email:    createEmployeeRequest.Email,
		Password: utils.HashPassword([]byte(createEmployeeRequest.Password)),
		Name:     createEmployeeRequest.Name,
		LastName: createEmployeeRequest.LastName,
		Phone:    createEmployeeRequest.Phone,
	}
	result := database.GormDB.Create(&employee)

	if result.Error != nil {
		return nil, result.Error
	}

	token, err := utils.GenerateEmployeeToken(employee.Name, employee.Email, "", "", employee.ID, 0, 0, "admin")
	if err != nil {
		return nil, err
	}

	createEmployeeResponse := modelsEmployees.CreateEmployeeResponse{
		ID:       employee.ID,
		Email:    employee.Email,
		Name:     employee.Name,
		LastName: employee.LastName,
		Token:    token,
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
		invitationToken, err := utils.GenerateInvitationToken(employeeTokenClaims.Email, employee.Email, employeeTokenClaims.ClinicID)
		if err != nil {
			return err
		}

		sendEmailManager := utils.SendEmailManager{
			ToEmail:         employee.Email,
			ToName:          employee.Name,
			FromName:        employeeTokenClaims.Name,
			ClinicName:      employeeTokenClaims.ClinicName,
			InvitationToken: invitationToken,
		}

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
	if deleteEmployeeClinic.EmployeeID != parentEmployeeID {
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
