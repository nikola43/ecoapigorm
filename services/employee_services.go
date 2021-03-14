package services

import (
	"errors"
	"fmt"
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

func Invite(employees []models.Employee) error {
	fmt.Println(employees)

	// validation ---------------------------------------------------------------------
	for _, employee := range employees {
		temp := models.Employee{}
		database.GormDB.Where("email = ?", employee.Email).Find(&employee)

		if temp.ID > 0 {
			// send link email
			sendEmailManager := utils.SendEmailManager{
				ToEmail:         employee.Email,
				ToName:          employee.Name,
				FromName:        employee.Email,
				CompanyName:     employee.Email,
				InvitationToken: employee.Email,
			}
			sendEmailManager.SendMail("invite_to_company.html", "Welcome")
		} else {
			// send signup email
			sendEmailManager := utils.SendEmailManager{
				ToEmail:         employee.Email,
				FromName:        employee.Email,
				CompanyName:     employee.Email,
				InvitationToken: employee.Email,
			}
			sendEmailManager.SendMail("invite_to_company.html", "Welcome")
		}
		fmt.Println(employee)
	}

	return nil
}
