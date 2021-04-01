package clients

type ChangePassClientRequest struct {
	ClientID int    `json:"client_id"`
	Token    string `json:"token"`
	Password string `json:"password"`
}
