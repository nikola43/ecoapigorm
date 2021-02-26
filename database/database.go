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

	database.GormDBResult := database.GormDB.Find(list)
	if database.GormDBResult.Error != nil {
		return nil, database.GormDBResult.Error
	}

	return list , nil
}

func GetClinicaById(id uint) (*models.Clinic, error) {
	var clinic *models.Clinic

	database.GormDBResult := database.GormDB.First(clinic,id)
	if database.GormDBResult.Error != nil {
		return nil, database.GormDBResult.Error
	}

	return clinic, nil
}
*/
