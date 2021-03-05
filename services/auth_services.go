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

	GormDBResult := database.GormDB.
		Where("email = ?", email).
		Find(&client)

	if GormDBResult.Error != nil {
		return &models.LoginClientResponse{}, GormDBResult.Error
	}

	match := utils.ComparePasswords(client.Password, []byte(password))
	if !match {
		return &models.LoginClientResponse{}, errors.New("not found")
	}

	token, err := utils.GenerateClientToken(client.Email, client.ID, client.ClinicID)
	if err != nil {
		return &models.LoginClientResponse{}, err
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

	GormDBResult := database.GormDB.
		Where("email = ?", email).
		Find(&employee)

	if GormDBResult.Error != nil {
		return nil, GormDBResult.Error
	}

	fmt.Println(employee)

	match := utils.ComparePasswords(employee.Password, []byte(password))
	if !match {
		return nil, errors.New("not found")
	}

	clinic:= models.Clinic{}
	database.GormDB.Model(&clinic).Select("clinics.id").Joins("left join employees on clinics.employee_id = employees.id").Where("employees.id = ?", employee.ID).Scan(&clinic)
	fmt.Println("clinic")
	fmt.Println(clinic.ID)

	// todo coger clinic id de la relacion
	token, err := utils.GenerateEmployeeToken(employee.Email, employee.ID, clinic.ID, employee.Role)
	if err != nil {
		return &models.LoginEmployeeResponse{}, err
	}

	clientEmployeeResponse := models.LoginEmployeeResponse{
		Id:       employee.ID,
		ClinicID: clinic.ID,
		Email:    employee.Email,
		Name:     employee.Name,
		Role:     employee.Role,
		LastName: employee.LastName,
		Token:    token,
	}

	return &clientEmployeeResponse, err
}
