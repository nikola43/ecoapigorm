package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	modelsClients "github.com/nikola43/ecoapigorm/models/clients"
	"github.com/nikola43/ecoapigorm/utils"
)

func CreateClient(createClientRequest *modelsClients.CreateClientRequest) (*modelsClients.CreateClientResponse, error) {
	//TODO validate

	client := models.Client{
		Email:    createClientRequest.Email,
		Password: utils.HashPassword([]byte(createClientRequest.Password)),
		Name:     createClientRequest.Name,
		LastName: createClientRequest.LastName,
	}
	result := database.GormDB.Create(&client)

	if result.Error != nil {
		return nil, result.Error
	}

	token, err := utils.GenerateClientToken(client.Email, client.ID, client.ClinicID)
	if err != nil {
		return nil, err
	}

	createClientResponse := modelsClients.CreateClientResponse{
		Id:       client.ID,
		Email:    client.Email,
		Name:     client.Name,
		LastName: client.LastName,
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

	return nil
}

func PassRecoveryClientService(request *modelsClients.PassRecoveryRequest) error {
	client := &models.Client{}

	GormDBResult := database.GormDB.
		Where("email = ?", request.Email).
		Find(&client)

	if GormDBResult.Error != nil {
		return GormDBResult.Error
	}

	apiTokenString, err := utils.GenerateClientToken(client.Email, client.ClinicID, client.ID)
	if err != nil {
		return err
	}

	recovery := models.Recovery{
		ClientID: client.ID,
		Token:    apiTokenString,
	}
	result := database.GormDB.Create(&recovery)
	if result.Error != nil {
		return result.Error
	}
	SendMailRecovery(client.Email, recovery.Token)

	return nil
}

func GetAllImagesByClientID(clientID string) ([]models.Image, error) {
	var list = make([]models.Image, 0)

	if err := database.GormDB.Find(&list).Where("client_id = ?", clientID).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func GetAllVideosByClientID(clientID string) ([]models.Video, error) {
	var list = make([]models.Video, 0)

	if err := database.GormDB.Find(&list).Where("client_id = ?", clientID).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func GetAllStreamingByClientID(clientID string) ([]models.Streaming, error) {
	var list = make([]models.Streaming, 0)

	if err := database.GormDB.Find(&list).Where("client_id = ?", clientID).Error; err != nil {
		return nil, err
	}

	return list, nil
}
