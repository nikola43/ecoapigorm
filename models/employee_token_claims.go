package models

type EmployeeTokenClaims struct {
	EmployeeID uint
	ClinicID   uint
	Email      string
	Role       string
	Exp        uint
}
