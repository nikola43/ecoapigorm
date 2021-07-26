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

func CreatePromoService(promoRequest *promos.CreatePromoRequest) (*models.Promo, error) {
	newPromo := &models.Promo{
		Title:    promoRequest.Title,
		Text:     promoRequest.Text,
		Week:     promoRequest.Week,
		ImageUrl: promoRequest.ImageUrl,
		StartAt:  promoRequest.StartAt,
		EndAt:    promoRequest.EndAt,
	}

	err := database.GormDB.Create(&newPromo).Error
	if err != nil {
		return nil, err
	}


	clinicPromo := new(models.ClinicPromo)
	clinicPromo.PromoID = newPromo.ID
	clinicPromo.ClinicID = promoRequest.ClinicID

	err = database.GormDB.Create(&clinicPromo).Error
	if err != nil {
		return nil, err
	}

	return newPromo, nil
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
