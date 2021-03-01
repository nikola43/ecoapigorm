package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	modelsClients "github.com/nikola43/ecoapigorm/models/clients"
	"github.com/nikola43/ecoapigorm/utils"
)

func CreateClient(newClient *modelsClients.CreateClientRequest) (*modelsClients.CreateClientResponse, error) {
	//TODO validate

	newClient.Password = utils.HashPassword([]byte(newClient.Password))
	result := database.GormDB.Create(newClient)

	if result.Error != nil {
		return nil, result.Error
	}

	token, err := utils.GenerateClientToken(newClient.Email)
	if err != nil {
		return nil, err
	}

	createClientResponse := modelsClients.CreateClientResponse{
		Name:     newClient.Name,
		LastName: newClient.LastName,
		Token:    token,
	}

	return &createClientResponse, result.Error
}

func ChangePassClientService(request *modelsClients.ChangePassClientRequest) error {
	client := &models.Client{}

	GormDBResult := database.GormDB.
		Find(&client, request.ClientId)

	if GormDBResult.Error != nil {
		return GormDBResult.Error
	}

	newPassHashed := utils.HashPassword([]byte(request.NewPass))

	database.GormDB.Model(&client).Update("password", newPassHashed)

	return  nil
}

func GetAllImagesByClientID(clientID string) ([]models.Image, error) {
	var list = make([]models.Image, 0)

	if err := database.GormDB.Find(&list).Where("id = ?", clientID).Error; err != nil {
		return nil, err
	}

	return list , nil
}
