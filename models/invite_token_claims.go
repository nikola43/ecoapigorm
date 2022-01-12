package models

type InviteTokenClaims struct {
	FromEmail string
	ToEmail   string
	ClinicID  uint
	Exp       uint
}
