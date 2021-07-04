package models

import (
	"github.com/nikola43/ecoapigorm/models/base"
)

type Company struct {
	base.CustomGormModel
	Name       string     `gorm:"type:varchar(32)" json:"name"`
	Clinics    []Clinic   `json:"clinics"`
}
