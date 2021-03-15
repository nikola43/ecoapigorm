package clients

type CreateCompanyRequest struct {
	Name string `json:"name" xml:"name" form:"name" validate:"required"`
}
