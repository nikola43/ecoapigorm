package clients

type CreateCompanyResponse struct {
	ID         uint   `json:"id"`
	EmployeeID uint   `json:"employee_id"`
	Name       string `json:"name"`
}
