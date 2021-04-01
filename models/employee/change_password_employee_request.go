package clients

type ChangePasswordEmployeeRequest struct {
	ID       uint   `json:"id"`
	Token    string `json:"token"`
	Password string `json:"password"`
}
