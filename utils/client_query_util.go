package utils

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
)

func GetClientByID(clientID uint) *models.Client {
	client := &models.Client{}

	err := database.GormDB.First(&client, clientID)
	if err != nil {
		return nil
	}

	if client.ID < 1 {
		return nil
	}
	return client
}

func GetClientByEmail(clientEmail string) *models.Client {
	client := &models.Client{}

	err := database.GormDB.Where("email = ?", clientEmail).Find(&client)
	if err != nil {
		return nil
	}

	if client.ID < 1 {
		return nil
	}
	return client
}
