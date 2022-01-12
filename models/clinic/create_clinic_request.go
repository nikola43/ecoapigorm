package clients

type CreateClinicRequest struct {
	CompanyID uint `json:"company_id" xml:"company_id" form:"company_id" validate:"required"`
	Name      string `json:"name" xml:"name" form:"name" validate:"required"`
}
