package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	models "github.com/nikola43/ecoapigorm/models"
)

func GetClinicas() ([]models.Clinic, error) {
	var list []models.Clinic

	GormDBResult := database.GormDB.Find(list)
	if GormDBResult.Error != nil {
		return nil, GormDBResult.Error
	}

	return list , nil
}

func GetClinicaById(id uint) (*models.Clinic, error) {
	var clinic *models.Clinic

	GormDBResult := database.GormDB.First(clinic,id)
	if GormDBResult.Error != nil {
		return nil, GormDBResult.Error
	}

	return clinic, nil
}
