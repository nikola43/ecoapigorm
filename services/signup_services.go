package services

import (
	"database/sql"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/utils"
)

func SignUpClient(name, email, password string) (string, error) {
	client := &models.Client{}
	GormDBResult := database.GormDB.
		Where("email = ?", email).
		Find(&client)

	if GormDBResult.Error != nil {
		return "", GormDBResult.Error
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
	GormDBResult := database.GormDB.
		Where("email = ?", email).
		Find(&client)

	if GormDBResult.Error != nil {
		return "", GormDBResult.Error
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
