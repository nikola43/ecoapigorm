package services

import (
	Database "github.com/nikola43/ecoapigorm/database"
	models "github.com/nikola43/ecoapigorm/models/responses"
)

func GetClinicas() ([]models.Clinic, error) {
	var list []models.Clinic

	dbResult := Database.DB.Find(list)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return list , nil
}

func GetClinicaById(id uint) (*models.Clinic, error) {
	var clinic *models.Clinic

	dbResult := Database.DB.First(clinic,id)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return clinic, nil
}