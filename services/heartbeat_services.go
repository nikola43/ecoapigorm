package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
)

func GetHeartbeatByClientID(clientID uint) (*models.Heartbeat, error) {
	var heartbeat = &models.Heartbeat{}

	if err := database.GormDB.Where("client_id = ?", clientID).First(&heartbeat).Error; err != nil {
		return nil, err
	}

	return heartbeat, nil
}

func GetHeartbeatByClientAndClinicID(clientID uint,clinicID uint) (*models.Heartbeat, error) {
	var heartbeat = &models.Heartbeat{}

	if err := database.GormDB.
		Where("client_id = ? AND clinic_id = ?", clientID,clinicID).
		First(&heartbeat).Error;
	err != nil {
		return nil, err
	}

	return heartbeat, nil
}
