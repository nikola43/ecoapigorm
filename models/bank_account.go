package models

import "github.com/nikola43/ecoapigorm/models/base"

type BankAccount struct {
	base.CustomGormModel
	PaymentMethodID uint   `json:"payment_method_id"`
	Iban            string `json:"iban"`
}
