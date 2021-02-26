package services

import (
	"database/sql"
	Database "github.com/nikola43/ecoapigorm/database"
	models "github.com/nikola43/ecoapigorm/models/responses"
	"github.com/nikola43/ecoapigorm/utils"
)

func SignUpClient(name, email, password string) (string, error) {
	client := &models.Client{}
	dbResult := Database.DB.
		Where("email = ?", email).
		Find(&client)

	if dbResult.Error != nil {
		return "", dbResult.Error
	}

	match := utils.ComparePasswords(client.Password, []byte(password))
	if match == false {
		return "", sql.ErrNoRows
	}

	// remove password
	//client.Password = ""

	token, err := utils.GenerateClientToken(client.Email)
	if err != nil {
		return "", err
	}
	return token, err
}

func SignUpEmployee(email, password string) (string, error) {
	client := &models.Client{}
	dbResult := Database.DB.
		Where("email = ?", email).
		Find(&client)

	if dbResult.Error != nil {
		return "", dbResult.Error
	}

	match := utils.ComparePasswords(client.Password, []byte(password))
	if match == false {
		return "", sql.ErrNoRows
	}

	// remove password
	//client.Password = ""

	token, err := utils.GenerateClientToken(client.Email)
	if err != nil {
		return "", err
	}
	return token, err
}
