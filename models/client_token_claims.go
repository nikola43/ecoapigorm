package models

type ClientTokenClaims struct {
	ID uint
	Email    string
	Exp      uint
}
