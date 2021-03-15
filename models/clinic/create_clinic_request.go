package clients

type CreateClinicRequest struct {
	EmployeeID uint `json:"employee_id"`
	Name       string `json:"name" validate:"required"`
}
