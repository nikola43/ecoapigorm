package models

type EmployeeTokenClaims struct {
	ID          uint
	ClinicID    uint
	Name        string
	CompanyID   uint
	CompanyName string
	ClinicName  string
	Email       string
	Role        string
	Exp         uint
}
