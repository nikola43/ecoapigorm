package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models/promos"
)

func GetAllPromos() ([]promos.Promo, error) {
	var list = make([]promos.Promo, 0)
	result := database.GormDB.
		Find(&list)

	if result.Error != nil {
		return nil, result.Error
	}

	return list, result.Error
}
