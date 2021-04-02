package services

import (
	"errors"
	"fmt"
	linq "github.com/ahmetb/go-linq/v3"
	database "github.com/nikola43/ecoapigorm/database"
	models "github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/clients"
	clinicModels "github.com/nikola43/ecoapigorm/models/clinic"
	_ "github.com/nikola43/ecoapigorm/models/employee"
	"github.com/nikola43/ecoapigorm/models/promos"
	"github.com/nikola43/ecoapigorm/models/streaming"
	"github.com/nikola43/ecoapigorm/utils"
	_ "github.com/nikola43/ecoapigorm/utils"
	"gorm.io/gorm"
)

func CreateClinic(companyID uint, createClinicRequest *clinicModels.CreateClinicRequest) (*clinicModels.CreateClinicResponse, error) {
	// give 30 credits first time
	var clinics = make([]*models.Clinic, 0)
	credits := 0

	GormDBResult := database.GormDB.Where("company_id = ?", companyID).Find(&clinics)
	if GormDBResult.Error != nil {

	}

	fmt.Println(clinics)
	fmt.Println(len(clinics))
	if len(clinics) == 0 {
		credits = 30
	}

	clinic := models.Clinic{
		Name:             createClinicRequest.Name,
		//EmployeeID:       employeeID,
		CompanyID: companyID,
		AvailableCredits: uint(credits),
	}
	result := database.GormDB.Create(&clinic)

	if result.Error != nil {
		return nil, result.Error
	}

	createEmployeeResponse := &clinicModels.CreateClinicResponse{
		ID:               clinic.ID,
		Name:             clinic.Name,
		//EmployeeID:       clinic.EmployeeID,
		CompanyID: companyID,
		AvailableCredits: uint(credits),
	}

	return createEmployeeResponse, result.Error
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
	listCLinicClients := make([]models.ClinicClient, 0)
	clinic := models.Clinic{}
	clientsList := make([]models.Client, 0)
	var totalSize uint = 0
	//videosSize := 0
	//heartbeatSize := 0

	database.GormDB.First(&clinic, id)
	database.GormDB.Where("clinic_id = ?", id).Find(&listCLinicClients)

	clientIds := make([]uint,0)
	linq.From(listCLinicClients).SelectT(func(u models.ClinicClient) uint { return u.ClientID }).ToSlice(&clientIds)

	database.GormDB.Find(&clientsList,&clientIds)

	for _,client := range clientsList {

		var size uint = 0
		totalSize = 0

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

		list = append(
			list,
			clients.ListClientResponse{
				ID:            client.ID,
				ClinicID:      clinic.ID,
				Email:         client.Email,
				Name:          client.Name,
				LastName:      client.LastName,
				Phone:         client.Phone,
				CreatedAt:     client.CreatedAt,
				PregnancyDate: client.PregnancyDate,
				UsedSize:      totalSize,
			},
		)
	}

/*	database.GormDB.Model(models.Client{}).Select(
		"distinct clients.id, "+
			"clinics.id as clinic_id, "+
			"clinics.name as clinic_name, "+
			"clients.email, "+
			"clients.name, "+
			"clients.last_name, "+
			"clients.phone, "+
			"calculators.week, "+
			"clients.created_at, "+
			"clients.pregnancy_date").

		Joins(
			"left join clinics "+
				"on clinics.id = clients.clinic_id").

		Joins("left join calculators "+
			"on clients.id = calculators.client_id").

		Where("clinics.id = ?", id).Scan(&list)

	for index, client := range list {
		var size uint = 0
		totalSize = 0

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

		list[index].UsedSize = totalSize
	}*/

	return list, nil
}

func CreateClientFromClinic(createClientRequest *clients.CreateClientRequest) (*clients.ListClientResponse, error) {
	client := models.Client{}
	clinic := models.Clinic{}
	//clinicOwnerParentEmployeeClinic := clinicModels.Clinic{}
	//useParentEmployeeClinicAvailableUsers := false
	useClinicAvailableUsers := false

	// check if client already exist
	if err := database.GormDB.
		Where("email = ?", createClientRequest.Email).
		Find(&client).Error; err != nil {
		return nil, err
	}

	if client.ID > 0 {
		return nil, errors.New("client already exist")
	}

	// check if client has been created by clinic

	if err := database.GormDB.First(&clinic, createClientRequest.ClinicID).Error; err != nil {
		return nil, errors.New("clinic_id not found")
	}

	// check if clinic has sufficient credits
	if clinic.AvailableCredits > 0 {
		useClinicAvailableUsers = true
	} else {
/*		// get clinic owner
		clinicOwnerEmployee := models.Employee{}
		if err := database.GormDB.First(&clinicOwnerEmployee, clinic.EmployeeID).Error; err != nil {
			return nil, errors.New("employee_id not found")
		}

		// check if has parent employee
		if clinicOwnerEmployee.ParentEmployeeID > 0 {
			// find parent employee
			clinicOwnerParentEmployee := models.Employee{}
			if err := database.GormDB.First(&clinicOwnerParentEmployee, clinicOwnerEmployee.ParentEmployeeID).Error; err != nil {
				return nil, errors.New("parent_employee_id not found")
			}

			// if find parent employee
			if clinicOwnerParentEmployee.ID > 0 {
				// get clinic owner employee clinic
				database.GormDB.Model(clinicModels.Clinic{}).Select(
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
					return nil, errors.New("parent employee not extends clients, insufficient credits")
				}
			} else {
				return nil, errors.New("parent_employee_id not found, insufficient credits")
			}
		} else {
			return nil, errors.New("insufficient credits")
		}*/
	}

	client = models.Client{
		Email:         createClientRequest.Email,
		Password:      utils.HashPassword([]byte("mimatrona")),
		Name:          createClientRequest.Name,
		LastName:      createClientRequest.LastName,
		PregnancyDate: createClientRequest.PregnancyDate,
		Phone:         createClientRequest.Phone,
	}
	result := database.GormDB.Create(&client)

	clinicClient := &models.ClinicClient{
		ClinicID:  clinic.ID,
		ClientID:  client.ID,
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

	// check if client has been created by clinic
	if useClinicAvailableUsers {
		database.GormDB.Model(&clinic).Update("available_credits", clinic.AvailableCredits-1)
	}

	// check if client has been created by clinic
/*	if useParentEmployeeClinicAvailableUsers {
		database.GormDB.Model(&clinicOwnerParentEmployeeClinic).Update(
			"available_credits", clinicOwnerParentEmployeeClinic.AvailableCredits-1)
	}*/

	return listClientResponse, result.Error
}

func GetAllPromosByClinicID(clinicID string) ([]promos.Promo, error) {
	var promos = make([]promos.Promo, 0)

	err := database.GormDB.Where("clinic_id = ?", clinicID).Find(&promos)
	if err.Error != nil {
		return nil, err.Error
	}

	return promos, nil
}

func GetAllStreamingByClinicID(clinicID string) ([]streaming.Streaming, error) {
	var list = make([]streaming.Streaming, 0)
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

	// check if client not is already linked by other clinic
	//todo
/*	if client.ClinicID > 0 {
		return errors.New("client is already linked by other clinic")
	}*/

	clinicClient := &models.ClinicClient{
		ClinicID:  clinic.ID,
		ClientID:  client.ID,
	}
	result = database.GormDB.Create(&clinicClient)

	/*//update clinic id
	//client.ClinicID = clinic.ID
	result = database.GormDB.Save(&client)*/
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteClinicByID(clinicID uint) error {
	deleteClinic := new(models.Clinic)
	clinicClients := make([]models.Client, 0)
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

func GetPromosByWeekAndClinicID(week, clinicID uint) ([]promos.Promo, error) {
	var list = make([]promos.Promo, 0)
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
