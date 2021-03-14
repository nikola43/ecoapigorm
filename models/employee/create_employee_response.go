package clients

type CreateEmployeeResponse struct {
	ID     uint `json:"id"`
	Email     string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Token    string `json:"token"`
}

