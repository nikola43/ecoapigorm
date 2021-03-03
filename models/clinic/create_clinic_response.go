package clients

type CreateClinicResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	EmployeeID uint `json:"employee_id"`
}
