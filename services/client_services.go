package services

import (
	"errors"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	modelsClients "github.com/nikola43/ecoapigorm/models/clients"
	"github.com/nikola43/ecoapigorm/models/streaming"
	"github.com/nikola43/ecoapigorm/utils"
)

func CreateClientFromApp(createClientRequest *modelsClients.CreateClientFromAppRequest) (*modelsClients.CreateClientResponse, error) {
	client := new(models.Client)

	// check if client already exist
	utils.GetModelByField(client, "email", createClientRequest.Email)
	if client.ID > 0 {
		return nil, errors.New("client already exist")
	}

	client = &models.Client{
		Email:         createClientRequest.Email,
		Password:      utils.HashPassword([]byte(createClientRequest.Password)),
		Name:          createClientRequest.Name,
		LastName:      createClientRequest.LastName,
		Phone:         createClientRequest.Phone,
		PregnancyDate: createClientRequest.PregnancyDate,
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
		ID:            client.ID,
		Email:         client.Email,
		Name:          client.Name,
		LastName:      client.LastName,
		Phone:         client.Phone,
		PregnancyDate: client.PregnancyDate,
		Token:         token,
	}

	return &createClientResponse, result.Error
}

func ChangePassClientService(request *modelsClients.ChangePassClientRequest) error {
	client := &models.Client{}

	GormDBResult := database.GormDB.
		First(&client, request.ID)

	if GormDBResult.Error != nil {
		return GormDBResult.Error
	}

	newPassHashed := utils.HashPassword([]byte(request.Password))

	database.GormDB.Model(&client).Update("password", newPassHashed)

	return nil
}

func UpdateClientService(id uint, updateClientRequest *modelsClients.UpdateClientRequest) (*models.Client, error) {
	client := &models.Client{}

	GormDBResult := database.GormDB.
		Find(&client, id)

	if GormDBResult.Error != nil {
		return nil, GormDBResult.Error
	}

	// database.GormDB.Model(&client).Update("pregnancy_date", nil)

	database.GormDB.Model(&client).Update("pregnancy_date", updateClientRequest.PregnancyDate)

	GormDBResult = database.GormDB.
		Model(&client).
		Updates(models.Client{Name: updateClientRequest.Name,
			LastName: updateClientRequest.LastName,
			Phone:    updateClientRequest.Phone})

	if GormDBResult.Error != nil {
		return nil, GormDBResult.Error
	}

	return client, nil
}

func GetClientByEmail(clientEmail string) (*models.Client, error) {
	client := &models.Client{}

	GormDBResult := database.GormDB.
		Where("email = ?", clientEmail).
		Find(&client)

	if GormDBResult.Error != nil {
		return nil, GormDBResult.Error
	}

	if client.ID < 1 {
		return nil, errors.New("not found")
	}

	return client, nil
}

func GetClientById(clientID uint) (*modelsClients.ListClientResponse, error) {
	client := &models.Client{}

	GormDBResult := database.GormDB.
		Where("id = ?", clientID).
		Find(&client)

	if GormDBResult.Error != nil {
		return nil, GormDBResult.Error
	}

	if client.ID < 1 {
		return nil, errors.New("client not found")
	}


	var size uint = 0
	var totalSize uint = 0

	// get images size
	database.GormDB.Table("images").Where("client_id = ?", client.ID).Select("IF(size IS NULL, 0, SUM(size)) as size").Scan(&size)
	totalSize += size

	//get videos size
	size = 0
	database.GormDB.Table("videos").Where("client_id = ?", client.ID).Select("IF(size IS NULL, 0, SUM(size)) as size").Scan(&size)
	totalSize += size

	//get heartbeat size
	size = 0
	database.GormDB.Table("heartbeats").Where("client_id = ?", client.ID).Select("IF(size IS NULL, 0, SUM(size)) as size").Scan(&size)
	totalSize += size

	clientResponse := &modelsClients.ListClientResponse{
		ID:            client.ID,
		ClinicID:      client.ID,
		Email:         client.Email,
		Name:          client.Name,
		LastName:      client.LastName,
		Phone:         client.Phone,
		CreatedAt:     client.CreatedAt,
		PregnancyDate: client.PregnancyDate,
		UsedSize:      totalSize,
	}


	return clientResponse, nil
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

	if err := database.GormDB.Where("client_id = ?", clientID).Find(&list).Error;
	err != nil {
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

func UnassignClientByID(clientID uint, clinicId uint) error {
	deleteClinicClient := new(models.ClinicClient)

	/*// todo check clinic is who make action
	// check if employee exist
	utils.GetModelByField(deleteClient, "id", clientID)
	if deleteClient.ID < 1 {
		return errors.New("client not found")
	}*/

	GormDBResult := database.GormDB.
		Where("clinic_id = ? AND client_id = ?", clinicId, clientID).
		Find(&deleteClinicClient)

	if GormDBResult.Error != nil {
		return GormDBResult.Error
	}

	database.GormDB.Unscoped().Delete(deleteClinicClient)

	/*
		// delete employee
		result := database.GormDB.Delete(deleteClient)
		if result.Error != nil {
			return result.Error
		}
	*/

	// todo remove content

	return nil
}

func RefreshClient(clientID uint) (*models.LoginClientResponse, error) {
	client := new(models.Client)

	utils.GetModelByField(client, "id", clientID)
	if client.ID < 1 {
		return nil, errors.New("client not found")
	}

	token, err := utils.GenerateClientToken(client.Email, client.ID,1/* client.ClinicID*/)
	if err != nil {
		return nil, err
	}

	clientLoginResponse := &models.LoginClientResponse{
		Id:            client.ID,
		Email:         client.Email,
		Name:          client.Name,
		Phone:         client.Phone,
		LastName:      client.LastName,
		Token:         token,
		ClinicID:      1,//client.ClinicID,
		PregnancyDate: client.PregnancyDate,
	}

	return clientLoginResponse, nil
}
