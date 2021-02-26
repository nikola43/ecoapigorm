package models

type ClientLoginResponse struct {
	ClinicID uint   `json:"clinic_id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Token    string `json:"token"`
}
