package clients

type CreateCompanyResponse struct {
	ID         uint   `json:"id"`
	OwnerEmployeeID uint   `json:"owner_employee_id"`
	Name       string `json:"name"`
	Token      string `json:"token"`
	CreatedAt  string `json:"created_at"`
}
