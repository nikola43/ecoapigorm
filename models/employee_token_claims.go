package models

type EmployeeTokenClaims struct {
	ID          uint
	ClinicID    uint
	Name        string
	CompanyName string
	ClinicName  string
	Email       string
	Role        string
	Exp         uint
}
