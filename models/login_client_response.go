package models

import (
	"github.com/nikola43/ecoapigorm/models/base"
)

type LoginClientResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Phone    string `json:"phone"`
	Token    string `json:"token"`
	ClinicID uint `json:"clinic_id"`
	PregnancyDate base.CustomNullTime `json:"pregnancy_date" xml:"pregnancy_date"`
}
