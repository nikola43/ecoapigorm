package services

import (
	"errors"
	"fmt"
	database "github.com/nikola43/ecoapigorm/database"
	models "github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/clients"
	clinicModels "github.com/nikola43/ecoapigorm/models/clinic"
	_ "github.com/nikola43/ecoapigorm/models/employee"
	"github.com/nikola43/ecoapigorm/models/promos"
	"github.com/nikola43/ecoapigorm/models/streaming"
	"github.com/nikola43/ecoapigorm/utils"
	_ "github.com/nikola43/ecoapigorm/utils"
)

func CreateClinic(createEmployeeRequest *clinicModels.CreateClinicRequest) (*clinicModels.CreateClinicResponse, error) {
	clinic := models.Clinic{
		Name:       createEmployeeRequest.Name,
		EmployeeID: createEmployeeRequest.EmployeeID,
	}
	result := database.GormDB.Create(&clinic)

	if result.Error != nil {
		return nil, result.Error
	}

	createEmployeeResponse := &clinicModels.CreateClinicResponse{
		ID:         clinic.ID,
		Name:       clinic.Name,
		EmployeeID: clinic.EmployeeID,
	}

	return createEmployeeResponse, result.Error
}

func GetClinicas() ([]models.Clinic, error) {
	var list []models.Clinic

	GormDBResult := database.GormDB.Find(list)
	if GormDBResult.Error != nil {
		return nil, GormDBResult.Error
	}

	return list, nil
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

	database.GormDB.Model(models.Client{}).Select(
		"distinct clients.id, "+
			"clinics.id as clinic_id, "+
			"clinics.name as clinic_name, "+
			"clients.email, "+
			"clients.name, "+
			"clients.last_name, "+
			"clients.phone, "+
			"calculators.week, "+
			"clients.created_at").

		Joins(
			"left join clinics "+
				"on clinics.id = clients.clinic_id").

		Joins("left join calculators "+
			"on clients.id = calculators.client_id").

		Where("clinics.id = ?", id).Scan(&list)
	return list, nil
}

func CreateClientFromClinic(createClientRequest *clients.CreateClientRequest) (*clients.ListClientResponse, error) {
	//TODO validate
	client := models.Client{}
	clinic := models.Clinic{}
	clinicOwnerParentEmployeeClinic := models.Clinic{}
	useParentEmployeeClinicAvailableUsers := false
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
		// get clinic owner
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
					return nil, errors.New("parent employee not extends clients, insufficient credits")
				}
			} else {
				return nil, errors.New("parent_employee_id not found, insufficient credits")
			}
		} else {
			return nil, errors.New("insufficient credits")
		}
	}

	client = models.Client{
		ClinicID: createClientRequest.ClinicID,
		Email:    createClientRequest.Email,
		Password: utils.HashPassword([]byte("mimatrona")),
		Name:     createClientRequest.Name,
		LastName: createClientRequest.LastName,
		Phone:    createClientRequest.Phone,
	}
	result := database.GormDB.Create(&client)

	if result.Error != nil {
		return nil, result.Error
	}

	listClientResponse := &clients.ListClientResponse{
		ID:        client.ID,
		ClinicID:  client.ClinicID,
		Email:     client.Email,
		Name:      client.Name,
		LastName:  client.LastName,
		Phone:     client.Phone,
		Week:      createClientRequest.Week,
		CreatedAt: client.CreatedAt.String(),
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
