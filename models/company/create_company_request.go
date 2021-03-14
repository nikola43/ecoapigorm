package clients

type CreateCompanyRequest struct {
	Name       string `json:"name" validate:"required"`
}
