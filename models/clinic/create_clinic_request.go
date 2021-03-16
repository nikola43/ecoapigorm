package clients

type CreateClinicRequest struct {
	Name string `json:"name" xml:"name" form:"name" validate:"required"`
}
