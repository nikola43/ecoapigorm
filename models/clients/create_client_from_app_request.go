package clients

import (
	"github.com/nikola43/ecoapigorm/models/base"
)


type CreateClientFromAppRequest struct {
	Name          string              `json:"name" validate:"required"`
	ClinicID      uint                `json:"clinic_id"`
	LastName      string              `json:"last_name" validate:"required"`
	Email         string              `json:"email" xml:"email" form:"email" validate:"required,email"`
	Password      string              `json:"password" xml:"password" form:"password" validate:"required"`
	PregnancyDate base.CustomNullTime `json:"pregnancy_date" xml:"pregnancy_date"`
	Phone         string              `json:"phone" xml:"phone" form:"phone"`
}
