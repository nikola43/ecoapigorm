package models

type ClientTokenClaims struct {
	ID uint
	ClinicID uint
	Email    string
	Exp      uint
}
