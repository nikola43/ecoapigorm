package clients

import "time"

type CreateClientResponse struct {
	ID       uint   `json:"id"`
	ClinicID uint   `json:"clinic_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	PregnancyDate    time.Time  `json:"pregnancy_date" xml:"pregnancy_date"`
	Token    string `json:"token"`
	Phone    string `json:"phone"`
}
