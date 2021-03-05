package clients

type CreateClientResponse struct {
	Id     uint `json:"id"`
	ClinicId     uint `json:"clinic_id"`
	Email     string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Token    string `json:"token"`
}

