package clients

type ChangePassClientRequest struct {
	ID       uint   `json:"id"`
	Token    string `json:"token"`
	Password string `json:"password"`
}
