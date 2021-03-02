package models

type PaymentMethod struct {
	CustomGormModel
	EmployeeID    uint `json:"employee_id"`
	CreditCardID  int `json:"credit_card_id"`
	BankAccountID int `json:"bank_account_id"`

	CreditCard  []CreditCard  `json:"credit_cards"`
	BankAccount []BankAccount `json:"bank_accounts"`
}
