package models

type BankAccount struct {
	CustomGormModel
	PaymentMethodID uint   `json:"payment_method_id"`
	Iban            string `json:"iban"`
}
