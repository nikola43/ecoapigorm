package models

type LoginClientResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Token    string `json:"token"`
	ClinicID uint `json:"clinic_id"`
}
