package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	models "github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/clients"
	_ "github.com/nikola43/ecoapigorm/models/employee"
	clinicModels "github.com/nikola43/ecoapigorm/models/clinic"
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

func GetClientsByClinicID(id uint) ([]clients.ListClientRequest, error) {
	list := make([]clients.ListClientRequest, 0)

	database.GormDB.Model(models.Client{}).Select(
		"clients.id, " +
			"clinics.id as clinic_id, " +
			"clinics.name as clinic_name, " +
			"clients.email, " +
			"clients.name, " +
			"clients.last_name, " +
			"clients.phone, " +
			"calculators.week, " +
			"clients.created_at").

		Joins(
			"inner join clinics " +
			"on clinics.id = clients.clinic_id").

		Joins("inner join calculators " +
			"on clients.id = calculators.client_id").

		Where("clinics.id = ?", id).Scan(&list)
	return list, nil
}
