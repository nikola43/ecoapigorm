package models

import (
	"github.com/nikola43/ecoapigorm/models/base"
)

type Company struct {
	base.CustomGormModel
	EmployeeID uint       `json:"employee_id"`
	OwnerEmployeeID uint       `json:"owner_employee_id"`
	Name       string     `gorm:"type:varchar(32)" json:"name"`
	Employees  []Employee `json:"employees"`
	Clinics  []Clinic `json:"clinics"`
}
