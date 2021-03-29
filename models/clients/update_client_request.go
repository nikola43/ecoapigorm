package clients

import "time"

type UpdateClientRequest struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Phone    string `json:"phone" xml:"phone" form:"phone"`
	PregnancyDate time.Time `json:"pregnancy_date" xml:"pregnancy_date"`
}
