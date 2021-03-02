package models

type ClientTokenClaims struct {
	ClientID uint
	ClinicID uint
	Email    string
	Exp     uint
}
