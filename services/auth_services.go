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

	if err = database.GormDB.Where("email = ?", email).Find(&client).Error; err != nil {
		return nil, err
	}

	if !utils.ComparePasswords(client.Password, []byte(password)) {
		return nil, errors.New("not found")
	}

	if token, err = utils.GenerateClientToken(client.Email, client.ID, client.ClinicID); err != nil {
		return nil, err
	}

	clientLoginResponse := models.LoginClientResponse{
		Id:       client.ID,
		Email:    client.Email,
		Name:     client.Name,
		LastName: client.LastName,
		Token:    token,
	}

	return &clientLoginResponse, err
}

func LoginEmployee(email, password string) (*models.LoginEmployeeResponse, error) {
	employee := &models.Employee{}
	clinic := models.Clinic{}
	company := models.Company{}
	token := ""
	var err error

	if err := database.GormDB.Where("email = ?", email).Find(&employee).Error; err != nil {
		return nil, err
	}

	match := utils.ComparePasswords(employee.Password, []byte(password))
	if !match {
		return nil, errors.New("not found")
	}

	database.GormDB.Model(&clinic).
		Joins("left join employees on clinics.employee_id = employees.id").
		Where("employees.id = ?", employee.ID).
		Find(&clinic)

	database.GormDB.Model(&company).
		Where("id = ?", employee.CompanyID).
		Find(&company)

	fmt.Println("clinicID")
	fmt.Println(clinic.ID)

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
		Clinic:       clinic,
	}

	return &clientEmployeeResponse, err
}
