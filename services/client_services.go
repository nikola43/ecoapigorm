package services

import (
	"errors"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	modelsClients "github.com/nikola43/ecoapigorm/models/clients"
	"github.com/nikola43/ecoapigorm/models/streaming"
	"github.com/nikola43/ecoapigorm/utils"
)

func CreateClient(createClientRequest *modelsClients.CreateClientFromAppRequest) (*modelsClients.CreateClientResponse, error) {
	client := new(models.Client)

	// check if client already exist
	utils.GetModelByField(client, "email", createClientRequest.Email)
	if client.ID > 0 {
		return nil, errors.New("client already exist")
	}

	client = &models.Client{
		Email:    createClientRequest.Email,
		Password: utils.HashPassword([]byte(createClientRequest.Password)),
		Name:     createClientRequest.Name,
		LastName: createClientRequest.LastName,
		Phone:    createClientRequest.Phone,
	}
	result := database.GormDB.Create(&client)

	if result.Error != nil {
		return nil, result.Error
	}

	token, err := utils.GenerateClientToken(client.Email, client.ID, 0)
	if err != nil {
		return nil, err
	}

	createClientResponse := modelsClients.CreateClientResponse{
		ID:       client.ID,
		Email:    client.Email,
		Name:     client.Name,
		LastName: client.LastName,
		Phone:    client.Phone,
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

func UpdateClientService(id uint ,request *modelsClients.UpdateClientRequest) error {
	client := &models.Client{}

	GormDBResult := database.GormDB.
		Find(&client, id)

	if GormDBResult.Error != nil {
		return GormDBResult.Error
	}

	GormDBResult = database.GormDB.
		Model(&client).
		Updates(models.Client{Name: request.Name, LastName: request.LastName, Phone: request.Phone})

	if GormDBResult.Error != nil {
		return GormDBResult.Error
	}

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

func GetClientById(clientID uint) (*models.Client, error) {
	client := &models.Client{}

	GormDBResult := database.GormDB.
		Where("id = ?", clientID).
		Find(&client)

	if GormDBResult.Error != nil {
		return nil, GormDBResult.Error
	}
	return client, nil
}

func GetAllImagesByClientID(clientID uint) ([]models.Image, error) {
	var list = make([]models.Image, 0)
	if err := database.GormDB.Where("client_id = ?", clientID).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func GetAllVideosByClientID(clientID uint) ([]models.Video, error) {
	var list = make([]models.Video, 0)

	if err := database.GormDB.Where("client_id = ?", clientID).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func GetAllHolographicsByClientID(clientID string) ([]models.Holographic, error) {
	var list = make([]models.Holographic, 0)

	if err := database.GormDB.Where("client_id = ?", clientID).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func GetAllStreamingByClientID(clientID string) ([]streaming.Streaming, error) {
	var list = make([]streaming.Streaming, 0)

	if err := database.GormDB.Where("client_id = ?", clientID).Find(&list).Error;
		err != nil {
		return nil, err
	}

	return list, nil
}

func DeleteClientByID(clientID uint) error {
	deleteClient := new(models.Client)

	// todo check clinic is who make action
	// check if employee exist
	utils.GetModelByField(deleteClient, "id", clientID)
	if deleteClient.ID < 1 {
		return errors.New("client not found")
	}

	// delete employee
	result := database.GormDB.Delete(deleteClient)
	if result.Error != nil {
		return result.Error
	}

	// todo remove content

	return nil
}
