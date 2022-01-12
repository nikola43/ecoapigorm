package services

import (
	errors "errors"
	"fmt"
	linq "github.com/ahmetb/go-linq/v3"
	database "github.com/nikola43/ecoapigorm/database"
	models "github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/clients"
	clinicModels "github.com/nikola43/ecoapigorm/models/clinic"
	_ "github.com/nikola43/ecoapigorm/models/employee"
	streamingModels "github.com/nikola43/ecoapigorm/models/streamings"
	"github.com/nikola43/ecoapigorm/utils"
	_ "github.com/nikola43/ecoapigorm/utils"
	"gorm.io/gorm"
	"sort"
)

func CreateClinic(createClinicRequest *clinicModels.CreateClinicRequest) (*clinicModels.CreateClinicResponse, error) {
	// give 30 credits first time
	var clinics = make([]*models.Clinic, 0)
	credits := 0

	GormDBResult := database.GormDB.Where("company_id = ?", createClinicRequest.CompanyID).Find(&clinics)
	if GormDBResult.Error != nil {

	}

	fmt.Println(clinics)
	fmt.Println(len(clinics))
	if len(clinics) == 0 {
		credits = 30
	}

	clinic := models.Clinic{
		Name:             createClinicRequest.Name,
		CompanyID:        createClinicRequest.CompanyID,
		AvailableCredits: uint(credits),
	}
	result := database.GormDB.Create(&clinic)

	if result.Error != nil {
		return nil, result.Error
	}

	createEmployeeResponse := &clinicModels.CreateClinicResponse{
		ID:               clinic.ID,
		Name:             clinic.Name,
		CompanyID:        createClinicRequest.CompanyID,
		AvailableCredits: uint(credits),
	}

	return createEmployeeResponse, result.Error
}

func GetCreditsClinicById(id uint) (uint, error) {
	clinic := models.Clinic{}

	if err := database.GormDB.First(&clinic, id).Error; err != nil {
		return 0, err
	}

	return clinic.AvailableCredits, nil
}

func GetClinicByID(id uint) (*models.Clinic, error) {
	clinic := models.Clinic{}

	if err := database.GormDB.First(&clinic, id).Error; err != nil {
		return nil, err
	}

	return &clinic, nil
}

func GetClientsByClinicID(id uint) ([]clients.ListClientResponse, error) {
	list := make([]clients.ListClientResponse, 0)
	listClinicClients := make([]models.ClinicClient, 0)
	clinic := models.Clinic{}
	clientsList := make([]models.Client, 0)

	database.GormDB.Find(&clinic, id)

	if clinic.ID < 1 {
		return nil, errors.New("clinic_id not found")
	}

	database.GormDB.Where("clinic_id = ?", id).Find(&listClinicClients)

	if len(listClinicClients) == 0 {
		return list, nil
	}

	clientIds := make([]uint, 0)
	linq.From(listClinicClients).
		SelectT(func(clinicClientRelation models.ClinicClient) uint { return clinicClientRelation.ClientID }).
		ToSlice(&clientIds)

	database.GormDB.Find(&clientsList, &clientIds)

	for i, client := range clientsList {
		totalSize := utils.CalculateTotalSizeByClient(client, clinic.ID)
		list = append(
			list,
			clients.ListClientResponse{
				ID:             client.ID,
				ClinicID:       clinic.ID,
				Email:          client.Email,
				Name:           client.Name,
				LastName:       client.LastName,
				Phone:          client.Phone,
				CreatedAt:      client.CreatedAt,
				PregnancyDate:  client.PregnancyDate,
				UsedSize:       totalSize,
				DiskQuoteLevel: listClinicClients[i].DiskQuoteLevel,
				ProfileColor:   client.ProfileColor,
			},
		)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[j].CreatedAt.Before(list[i].CreatedAt)
	})

	return list, nil
}

func CreateClientFromClinic(createClientRequest *clients.CreateClientRequest) (*clients.ListClientResponse, error) {
	client := models.Client{}
	clinic := models.Clinic{}

	// check if client already exist
	if err := database.GormDB.
		Where("email = ?", createClientRequest.Email).
		Find(&client).Error; err != nil {
		return nil, err
	}

	// check if client has been created by clinic
	if err := database.GormDB.First(&clinic, createClientRequest.ClinicID).Error; err != nil {
		return nil, errors.New("clinic_id not found")
	}

	clinicClient := &models.ClinicClient{}
	result := database.GormDB.Where("clinic_id = ? AND client_id", clinic.ID, client.ID).First(&clinicClient)

	if clinicClient.ClinicID == clinic.ID && clinicClient.ClientID == client.ID {
		return nil, errors.New("client already assigned to this clinic")
	}

	client = models.Client{
		Email:         createClientRequest.Email,
		Password:      utils.HashPassword([]byte("ecox5d")),
		Name:          createClientRequest.Name,
		LastName:      createClientRequest.LastName,
		PregnancyDate: createClientRequest.PregnancyDate,
		Phone:         createClientRequest.Phone,
		ProfileColor:  createClientRequest.ProfileColor,
	}
	result = database.GormDB.Create(&client)

	clinicClient = &models.ClinicClient{
		ClinicID: clinic.ID,
		ClientID: client.ID,
	}

	result = database.GormDB.Create(&clinicClient)
	if result.Error != nil {
		return nil, result.Error
	}

	listClientResponse := &clients.ListClientResponse{
		ID:            client.ID,
		ClinicID:      clinic.ID,
		Email:         client.Email,
		Name:          client.Name,
		LastName:      client.LastName,
		Phone:         client.Phone,
		PregnancyDate: createClientRequest.PregnancyDate,
		CreatedAt:     client.CreatedAt,
	}

	database.GormDB.Model(&clinic).Update("available_credits", clinic.AvailableCredits-1)
	return listClientResponse, result.Error
}

func GetAllPromosByClinicID(clinicID string) ([]*models.Promo, error) {
	clinic := new(models.Clinic)

	err := database.GormDB.
		Preload("Promos").
		First(&clinic, clinicID).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return clinic.Promos, nil
}

func GetAllPromosForClient(clientId uint, clinicId uint) ([]*models.Promo, error) {
	clinic := new(models.Clinic)

	/*
		clinicClient := make([]models.ClinicClient, 0)

		err := database.GormDB.
			Where("client_id = ? AND clinic_id = ?", clientId, clinicId).
			Find(&clinicClient)
		if err.Error != nil {
			return nil, err.Error
		}

		if len(clinicClient) == 0 {
			return promos, nil
		}

		// TODO Añadir week

	*/

	err := database.GormDB.
		Preload("Promos").
		First(clinic, clinicId)
	if err.Error != nil {
		return nil, err.Error
	}

	return clinic.Promos, nil
}

func GetAllStreamingByClinicID(clinicID string) ([]streamingModels.Streaming, error) {
	var list = make([]streamingModels.Streaming, 0)
	var clients = make([]models.Client, 0)
	var clientsIds = make([]uint, 0)

	// todo consultar solo id
	err := database.GormDB.Where("clinic_id = ?", clinicID).Find(&clients)
	if err.Error != nil {
		return nil, err.Error
	}

	for i := 0; i < len(clients); i++ {
		clientsIds = append(clientsIds, clients[i].ID)
	}

	fmt.Println("ss")
	if err := database.GormDB.Where("client_id IN (?)", clientsIds).Find(&list).Error;
		err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return list, nil
}

func GetEmployeesByClinicID(clinicID uint) ([]models.Employee, error) {
	list := make([]models.Employee, 0)

	if err := database.GormDB.Where("clinic_id = ?", clinicID).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func UpdateClinic(clinic *models.Clinic) (*models.Clinic, error) {
	findClinic := &models.Clinic{}

	result := database.GormDB.Where("id = ?", clinic.ID).First(&findClinic)
	if result.Error != nil {
		return nil, result.Error
	}

	result = database.GormDB.Save(&clinic)
	if result.Error != nil {
		return nil, result.Error
	}
	return clinic, nil
}

func LinkClient(clientID uint, clinicID uint) error {

	// todo check credits

	client := &models.Client{}
	clinic := &models.Clinic{}

	// check if exist client
	result := database.GormDB.Where("id = ?", clientID).First(&client)
	if result.Error != nil {
		return result.Error
	}

	// check if exist clinic
	result = database.GormDB.Where("id = ?", clinicID).First(&clinic)
	if result.Error != nil {
		return result.Error
	}

	clinicClient := &models.ClinicClient{}
	result = database.GormDB.Where("clinic_id = ? AND client_id = ?", clinicID, clientID).First(&clinicClient)
	if result.Error != nil {
		// todo revisar
		// return result.Error
	}

	// check if client not is already linked by other clinic
	if clinicClient.ClinicID > 0 && clinicClient.ClientID > 0 {
		return errors.New("client is already linked")
	}

	clinicClient = &models.ClinicClient{
		ClinicID: clinic.ID,
		ClientID: client.ID,
	}

	result = database.GormDB.Create(&clinicClient)
	if result.Error != nil {
		return result.Error
	}

	database.GormDB.Model(&clinic).Update("available_credits", clinic.AvailableCredits-1)

	return nil
}

func DeleteClinicByID(clinicID uint) error {
	deleteClinic := new(models.Clinic)
	clinicClients := make([]models.ClinicClient, 0)
	clinicEmployees := make([]models.Employee, 0)

	// todo check clinic is who make action
	// check if employee exist
	utils.GetModelByField(deleteClinic, "id", clinicID)
	if deleteClinic.ID < 1 {
		return errors.New("clinic not found")
	}

	err := database.GormDB.Where("clinic_id = ?", clinicID).Find(&clinicClients)
	if err.Error != nil {
		return err.Error
	}

	if len(clinicClients) > 0 {
		return errors.New("clinic has clients")
	}

	err = database.GormDB.Where("clinic_id = ?", clinicID).Find(&clinicEmployees)
	if err.Error != nil {
		return err.Error
	}

	if len(clinicEmployees) > 0 {
		return errors.New("clinic has employees")
	}

	// delete clinic
	result := database.GormDB.Delete(deleteClinic)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetPromosByWeekAndClinicID(week, clinicID uint) ([]models.Promo, error) {
	var list = make([]models.Promo, 0)
	var result *gorm.DB

	if week >= 1 && week <= 22 {
		result = database.GormDB.Where("clinic_id = ? AND week BETWEEN 1 AND 22", clinicID, week).Find(&list)
	} else if week >= 23 && week <= 25 {
		result = database.GormDB.Where("clinic_id = ? AND week BETWEEN 23 AND 25", clinicID, week).Find(&list)
	} else if week >= 26 && week <= 32 {
		result = database.GormDB.Where("clinic_id = ? AND week BETWEEN 26 AND 32", clinicID, week).Find(&list)
	} else if week >= 33 && week <= 40 {
		result = database.GormDB.Where("clinic_id = ? AND week BETWEEN 33 AND 40", clinicID, week).Find(&list)
	} else if week == 41 {
		result = database.GormDB.Where("clinic_id = ?", clinicID, week).Find(&list)
	}

	if result != nil && result.Error != nil {
		return nil, result.Error
	}

	return list, nil
}
