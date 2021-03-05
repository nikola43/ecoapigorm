package clients

type CreateCompanyRequest struct {
	EmployeeID uint `json:"employee_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
}
