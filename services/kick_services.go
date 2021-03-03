package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models/kicks"
)

func CreateNewKickService(kickRequest kicks.AddKickRequest) (kicks.Kick, error) {
	kick := kicks.Kick{
		Date:     kickRequest.Date,
		ClientId: kickRequest.ClientId,
	}
	result := database.GormDB.Create(&kick)

	if result.Error != nil {
		return kick, result.Error
	}

	return kick, result.Error
}

func DeleteKickByIdService(ClientId uint, kickId uint) error {
	result := database.GormDB.
		Where("client_id = ?", ClientId).
		Delete(kicks.Kick{}, kickId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAllKicksByClientService(clientID uint) ([]kicks.Kick, error) {
	var list = make([]kicks.Kick, 0)
	result := database.GormDB.
		Where("client_id = ?", clientID).
		Find(&list)

	if result.Error != nil {
		return nil, result.Error
	}

	return list, result.Error
}

func ResetAllKicksByClientService(clientID uint) error {
	result := database.GormDB.
		Where("client_id = ?", clientID).
		Delete(kicks.Kick{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
