package clients

type PassRecoveryRequest struct {
	Email    string `json:"email" xml:"email" form:"email"`
}