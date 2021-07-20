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

func CreatePromoService(promoRequest *promos.CreatePromoRequest) (*promos.Promo, error) {
	newPromo := promos.Promo{
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

	return &newPromo, nil
	clinicPromo := new(models.ClinicPromo)
	clinicPromo.PromoID = promo.ID
	clinicPromo.ClinicID = createPromoRequest.ClinicID

	err = database.GormDB.Create(&clinicPromo).Error
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
