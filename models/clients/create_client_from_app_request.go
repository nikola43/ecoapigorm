package clients

// todo crear otro modelo para el panel
type CreateClientFromAppRequest struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Email    string `json:"email" xml:"email" form:"email" validate:"required,email"`
	Password string `json:"password" xml:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" xml:"phone" form:"phone"`
}
