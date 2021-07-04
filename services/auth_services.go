package services

import (
	"errors"
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
		Id:       client.ID,
		Email:    client.Email,
		Name:     client.Name,
		Phone:    client.Phone,
		LastName: client.LastName,
		Token:    token,
		Clinics: client.Clinics,
		PregnancyDate: client.PregnancyDate,
	}

	return clientLoginResponse, err
}

func LoginEmployee(email, password string) (*models.LoginEmployeeResponse, error) {
	employee := new(models.Employee)
	clinic := new(models.Clinic)
	company := new(models.Company)
	var token string

	err := database.GormDB.Where("email = ?", email).Find(&employee).Error
	if err != nil {
		return nil, err
	}

	match := utils.ComparePasswords(employee.Password, []byte(password))
	if !match {
		return nil, errors.New("not found")
	}

	err = database.GormDB.First(&company,employee.CompanyID).Error
	if err != nil {
		return nil, err
	}

	//Si es admin le metemos todas las clinicas de la company
	if employee.Role == "admin" {
		database.GormDB.Where("company_id = ?", employee.CompanyID).First(&clinic) //TODO enviar todas en lugar de una
	}else {
		err = database.GormDB.First(&clinic, employee.ClinicID).Error
		if err != nil {
			return nil, err
		}
	}

	token, err = utils.GenerateEmployeeToken(
		employee.Name,
		company.ID,
		clinic.ID,
		employee.ID,
		employee.Email,
		company.Name,
		clinic.Name,
		employee.Role)
	if err != nil {
		return nil, err
	}

	clientEmployeeResponse := models.LoginEmployeeResponse{
		ID:           employee.ID,
		CompanyID:    employee.CompanyID,
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
