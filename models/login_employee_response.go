package models

type LoginEmployeeResponse struct {
	ID           uint   `json:"id"`
	CompanyID    uint   `json:"company_id"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	Name         string `json:"name"`
	IsFirstLogin bool   `json:"is_first_login"`
	LastName     string `json:"lastname"`
	Token        string `json:"token"`
	Clinic       Clinic `json:"clinic"`
}
