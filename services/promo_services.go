package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/promos"
)

func GetPromoByID(id uint) (*models.Promo, error) {
	promo := new(models.Promo)

	err := database.GormDB.First(&promo, id).Error
	if err != nil {
		return nil, err
	}

	return promo, nil
}

func CreatePromoService(createPromoRequest *promos.CreatePromoRequest) (*models.Promo, error) {
	promo := &models.Promo{
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
	promo := new(models.Promo)

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
