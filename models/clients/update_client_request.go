package clients

import (
	"github.com/nikola43/ecoapigorm/models/base"
)

type UpdateClientRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Phone    string `json:"phone" xml:"phone" form:"phone"`
	PregnancyDate base.CustomNullTime `json:"pregnancy_date" xml:"pregnancy_date"`
}
