package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models/promos"
)

func GetPromoByID(id uint) (*promos.Promo, error) {
	promo := new(promos.Promo)

	err := database.GormDB.First(&promo, id).Error
	if err != nil {
		return nil, err
	}

	return promo, nil
}

func CreatePromoService(createPromoRequest *promos.CreatePromoRequest) (*promos.Promo, error) {
	promo := &promos.Promo{
		ClinicID: createPromoRequest.ClinicID,
		Title:    createPromoRequest.Title,
		Text:     createPromoRequest.Text,
		Week:     createPromoRequest.Week,
		StartAt:  createPromoRequest.StartAt,
		EndAt:    createPromoRequest.EndAt,
	}

	err := database.GormDB.Create(&promo).Error
	if err != nil {
		return nil, err
	}

	return promo, nil
}

func DeletePromoByID(id uint) error {
	promo := new(promos.Promo)

	err := database.GormDB.First(&promo, id).Error
	if err != nil {
		return err
	}

	err = database.GormDB.Delete(promo).Error
	if err != nil {
		return err
	}

	return nil
}
