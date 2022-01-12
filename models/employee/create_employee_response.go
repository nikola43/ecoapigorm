package clients

type CreateEmployeeResponse struct {
	ID           uint   `json:"id"`
	ClinicID     uint   `json:"clinic_id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	LastName     string `json:"lastname"`
	Role         string `json:"role"`
}
