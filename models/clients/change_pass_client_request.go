package clients

type ChangePassClientRequest struct {
	ID       uint   `json:"id" validate:"required"`
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required"`
}
