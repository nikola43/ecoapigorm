package models

import "github.com/nikola43/ecoapigorm/models/base"

type Payment struct {
	base.CustomGormModel
	SessionID  uint    `json:"session_id"`
	EmployeeID uint    `gorm:"type:INTEGER not null" json:"employee_id"`
	ClinicID   uint    `gorm:"type:INTEGER null" json:"clinic_id"`
	Amount     float32 `json:"amount"`
	Quantity   uint    `json:"quantity"`
}
