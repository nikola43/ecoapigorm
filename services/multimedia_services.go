package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
)

func DeleteImage(imageID uint) error {
	image := &models.Image{}

	result := database.GormDB.First(&image, imageID)
	if result.Error != nil {
		return result.Error
	}

	// delete from S3

	// delete from DB
	result = database.GormDB.Unscoped().Delete(image)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
