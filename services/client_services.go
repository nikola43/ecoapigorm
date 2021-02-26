package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
)

func GetAllImagesByClientID(clientID string) ([]models.Image, error) {
	var list = make([]models.Image, 0)

	if err := database.GormDB.Find(&list).Where("id = ?", clientID).Error; err != nil {
		return nil, err
	}

	return list , nil
}
