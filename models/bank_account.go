package models

import (
	"gorm.io/gorm"
)

type BankAccount struct {
	gorm.Model
	PaymentMethodID uint   `json:"payment_method_id"`
	Iban            string `json:"iban"`
}
