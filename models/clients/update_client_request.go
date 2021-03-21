package clients

type UpdateClientRequest struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Phone    string `json:"phone" xml:"phone" form:"phone"`
}
