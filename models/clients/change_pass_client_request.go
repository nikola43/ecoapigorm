package clients

type ChangePassClientRequest struct {
	ClientId int `json:"clientid"`
	NewPass string `json:"newpass"`
}
