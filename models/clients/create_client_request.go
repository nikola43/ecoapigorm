package clients

type CreateClientRequest struct {
	ClinicID uint   `json:"clinic_id"`
	Name     string `json:"name" validate:"required"`
	LastName string `json:"lastname" validate:"required"`
	Email    string `json:"email" xml:"email" form:"email" validate:"required,email"`
	Password string `json:"password" xml:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" xml:"phone" form:"phone"`
}
