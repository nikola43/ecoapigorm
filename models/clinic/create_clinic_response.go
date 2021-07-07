package clients

type CreateClinicResponse struct {
	ID               uint   `json:"id"`
	EmployeeID       uint   `json:"employee_id"` //TODO revisar si hace falta
	CompanyID        uint   `json:"company_id"`
	Name             string `json:"name"`
	AvailableCredits uint   `json:"available_credits"`
}
