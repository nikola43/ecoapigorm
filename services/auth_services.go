package services

import (
	"errors"
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
