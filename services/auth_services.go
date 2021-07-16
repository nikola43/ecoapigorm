package services

import (
	"errors"
	"fmt"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/utils"
)

func LoginClient(email, password string) (*models.LoginClientResponse, error) {
	client := &models.Client{}
	var err error
	token := ""

	err = database.GormDB.
		Where("email = ?", email).
		Preload("Clinics").
		Find(&client).Error
	if err != nil {
		return nil, err
	}

	if !utils.ComparePasswords(client.Password, []byte(password)) {
		return nil, errors.New("not found")
	}

	if token, err = utils.GenerateClientToken(client.Email, client.ID); err != nil {
		return nil, err
	}

	clientLoginResponse := &models.LoginClientResponse{
		Id:            client.ID,
		Email:         client.Email,
		Name:          client.Name,
		Phone:         client.Phone,
		LastName:      client.LastName,
		Token:         token,
		Clinics:       client.Clinics,
		PregnancyDate: client.PregnancyDate,
	}

	return clientLoginResponse, err
}

func LoginEmployee(loginEmployeeRequest *models.LoginEmployeeRequest) (*models.LoginEmployeeResponse, error) {
	employee := new(models.Employee)
	clinic := new(models.Clinic)
	company := new(models.Company)
	var token string

	err := database.GormDB.Where("email = ?", loginEmployeeRequest.Email).Find(&employee).Error
	if err != nil {
		return nil, err
	}

	fmt.Println(employee)

	if employee.ID < 1 {
		return nil, errors.New("not found")
	}

	match := utils.ComparePasswords(employee.Password, []byte(loginEmployeeRequest.Password))
	if !match {
		return nil, errors.New("password not match")
	}

	err = database.GormDB.Where("owner_employee_id = ?", employee.ID).Find(&company).Error
	if err != nil {
		fmt.Println(company)
		// return nil, err
	}

	err = database.GormDB.First(&clinic, employee.ClinicID).Error
	if err != nil {
		fmt.Println(clinic)
		//return nil, err
	}

	token, err = utils.GenerateEmployeeToken(
		employee.Name,
		company.ID,
		clinic.ID,
		employee.ID,
		employee.Email,
		company.Name,
		employee.Role)
	if err != nil {
		return nil, err
	}

	clientEmployeeResponse := models.LoginEmployeeResponse{
		ID:           employee.ID,
		CompanyID:    company.ID,
		Email:        employee.Email,
		Name:         employee.Name,
		Role:         employee.Role,
		IsFirstLogin: employee.IsFirstLogin,
		LastName:     employee.LastName,
		Token:        token,
		Clinic:       *clinic,
	}

	return &clientEmployeeResponse, err
}
