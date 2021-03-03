package clients

type CreateClinicRequest struct {
	Name       string `json:"name" validate:"required"`
	EmployeeID uint `json:"employee_id" validate:"required"`
}
