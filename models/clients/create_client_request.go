package clients

type CreateClientRequest struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email" xml:"email" form:"email"`
	Password string `json:"password" xml:"password" form:"password"`
}