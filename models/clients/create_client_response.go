package clients

type CreateClientResponse struct {
	ID       uint   `json:"id"`
	ClinicID uint   `json:"clinic_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Token    string `json:"token"`
}
