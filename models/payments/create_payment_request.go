package payments

type CreatePaymentRequest struct {
	CompanyID  uint    `gorm:"type:INTEGER not null" json:"company_id"`
	ClinicID   uint    `gorm:"type:INTEGER null" json:"clinic_id"`
	EmployeeID uint    `gorm:"type:INTEGER not null" json:"employee_id"`
	Amount     float32 `json:"amount"`
	Quantity   uint    `json:"quantity"`
}
