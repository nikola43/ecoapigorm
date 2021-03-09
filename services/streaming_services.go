package services

import (
	"errors"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
)

func GetStreamingByCode(code string) (*models.Streaming, error) {
	streaming := &models.Streaming{}
	result := database.GormDB.
		Where("code = ?", code).
		Find(&streaming)

	if result.Error != nil {
		return nil, result.Error
	}

	if streaming.ID < 1 {
		return nil, errors.New("not found")
	}

	return streaming, nil
}
