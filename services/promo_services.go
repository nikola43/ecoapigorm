package services

import (
	"fmt"
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

func CreatePromoService(promoRequest promos.CreatePromoRequest) (*promos.Promo, error) {
	newPromo := promos.Promo{
		ClinicID:        promoRequest.ClinicID,
		Title:           promoRequest.Title,
		Text:            promoRequest.Text,
		ImageUrl:        promoRequest.ImageUrl,
		Week:            promoRequest.Week,
		StartAt:         promoRequest.StartAt,
		EndAt:           promoRequest.EndAt,
	}

	fmt.Println("CreatePromoRequest")
	fmt.Println(promoRequest)

	result := database.GormDB.Create(&newPromo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &newPromo, result.Error
}
