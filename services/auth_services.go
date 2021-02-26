package services

import (
	"database/sql"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/utils"
)

func LoginClient(email, password string) (*models.ClientLoginResponse, error) {
	client := &models.Client{}

	GormDBResult := database.GormDB.
		Where("email = ?", email).
		Find(&client)

	if GormDBResult.Error != nil {
		return &models.ClientLoginResponse{}, GormDBResult.Error
	}

	match := utils.ComparePasswords(client.Password, []byte(password))
	if match == false {
		return &models.ClientLoginResponse{}, sql.ErrNoRows
	}

	// remove password
	//client.Password = ""

	token, err := utils.GenerateClientToken(client.Email)
	if err != nil {
		return &models.ClientLoginResponse{}, err
	}
	clientLoginResponse := models.ClientLoginResponse{client.ClinicID, client.Name, client.LastName, token}

	return &clientLoginResponse, err
}
