package models

import "github.com/nikola43/ecoapigorm/models/base"

type CreditCard struct {
	base.CustomGormModel
	PaymentMethodID uint   `json:"payment_method_id"`
	LastFourNumbers uint   `json:"last_four_numbers"`
	Token           string `json:"token"`
}
