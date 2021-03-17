package services

import (
	"errors"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	modelsClients "github.com/nikola43/ecoapigorm/models/clients"
	"github.com/nikola43/ecoapigorm/models/streaming"
	"github.com/nikola43/ecoapigorm/utils"
)

func CreateClient(createClientRequest *modelsClients.CreateClientRequest) (*modelsClients.CreateClientResponse, error) {
	//TODO validate
	client := new(models.Client)
	clinic := new(models.Clinic)
	clinicOwnerParentEmployeeClinic := new(models.Clinic)
	useParentEmployeeClinicAvailableUsers := false
	useClinicAvailableUsers := false

	// check if client already exist
	utils.GetModelByField(client, "email", createClientRequest.Email)
	if client.ID > 0 {
		return nil, errors.New("client already exist")
	}

	// find clinic
	utils.GetModelByField(clinic, "id", createClientRequest.ClinicID)
	if clinic.ID < 1 {
		return nil, errors.New("clinic not found")
	}

	// check if clinic has sufficient credits
	if clinic.AvailableCredits > 0 {
		useClinicAvailableUsers = true
	} else {
		// get clinic owner
		clinicOwnerEmployee := new(models.Employee)
		utils.GetModelByField(clinicOwnerEmployee, "id", clinic.EmployeeID)
		if clinicOwnerEmployee == nil {
			return nil, errors.New("employee_id not found")
		}

		// check if has parent employee
		if clinicOwnerEmployee.ParentEmployeeID > 0 {
			// find parent employee
			clinicOwnerParentEmployee := models.Employee{}
			utils.GetModelByField(clinicOwnerParentEmployee, "id", clinicOwnerEmployee.ParentEmployeeID)
			if clinicOwnerEmployee == nil {
				return nil, errors.New("parent_employee_id not found")
			}

			// if find parent employee
			if clinicOwnerParentEmployee.ID > 0 {
				// get clinic owner employee clinic
				database.GormDB.Model(models.Clinic{}).Select(
					"clinics.id, clinics.extend_clients, clinics.available_clients").Joins(
					"inner join employees on clinics.employee_id = employees.id").Where(
					"employees.id = ?", clinicOwnerParentEmployee.ID).Scan(&clinicOwnerParentEmployeeClinic)

				if clinicOwnerParentEmployeeClinic.ExtendCredits {
					if clinicOwnerParentEmployeeClinic.AvailableCredits > 0 {
						useParentEmployeeClinicAvailableUsers = true
					} else {
						return nil, errors.New("insufficient parent employee credits")
					}
				} else {
					return nil, errors.New("parent employee not extends credits, insufficient credits")
				}
			} else {
				return nil, errors.New("parent_employee_id not found, insufficient credits")
			}
		} else {
			return nil, errors.New("insufficient credits")
		}
	}

	client = &models.Client{
		ClinicID: createClientRequest.ClinicID,
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

	token, err := utils.GenerateClientToken(client.Email, client.ID, client.ClinicID)
	if err != nil {
		return nil, err
	}

	createClientResponse := modelsClients.CreateClientResponse{
		ID:       client.ID,
		ClinicID: client.ClinicID,
		Email:    client.Email,
		Name:     client.Name,
		LastName: client.LastName,
		Token:    token,
	}

	// check if client has been created by clinic
	if useClinicAvailableUsers {
		database.GormDB.Model(&clinic).Update("available_credits", clinic.AvailableCredits-1)
	}

	// check if client has been created by clinic
	if useParentEmployeeClinicAvailableUsers {
		database.GormDB.Model(&clinicOwnerParentEmployeeClinic).Update(
			"available_credits", clinicOwnerParentEmployeeClinic.AvailableCredits-1)
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
