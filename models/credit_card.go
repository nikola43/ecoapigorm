package models

import (
	"gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	PaymentMethodID uint   `json:"payment_method_id"`
	LastFourNumbers uint   `json:"last_four_numbers"`
	Token           string `json:"token"`
}
