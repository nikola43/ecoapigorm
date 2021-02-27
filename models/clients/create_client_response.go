package clients

type CreateClientResponse struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Token    string `json:"token"`
}

