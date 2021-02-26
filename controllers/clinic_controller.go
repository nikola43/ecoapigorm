package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllClinics (context *fiber.Ctx) error {
	return context.SendString("Hello, World!")

}

/*
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
*/
