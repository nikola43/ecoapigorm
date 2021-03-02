package models

type CreditCard struct {
	CustomGormModel
	PaymentMethodID uint   `json:"payment_method_id"`
	LastFourNumbers uint   `json:"last_four_numbers"`
	Token           string `json:"token"`
}
