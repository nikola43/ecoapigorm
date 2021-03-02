package models

type LoginEmployeeResponse struct {
	Id       uint   `json:"id"`
	ClinicID uint   `json:"clinic_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Token    string `json:"token"`
}
