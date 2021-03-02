package services

import (
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

	token, err := utils.GenerateEmployeeToken(employee.Email, employee.ID, 0, "admin")
	if err != nil {
		return nil, err
	}

	createEmployeeResponse := modelsEmployees.CreateEmployeeResponse{
		Id:       employee.ID,
		Email:    employee.Email,
		Name:     employee.Name,
		LastName: employee.LastName,
		Token:    token,
	}

	return &createEmployeeResponse, result.Error
}
