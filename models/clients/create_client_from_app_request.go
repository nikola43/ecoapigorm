package clients

import "time"


type CreateClientFromAppRequest struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Email    string `json:"email" xml:"email" form:"email" validate:"required,email"`
	Password string `json:"password" xml:"password" form:"password" validate:"required"`
	PregnancyDate    time.Time   `json:"pregnancy_date" xml:"pregnancy_date"`
	Phone    string `json:"phone" xml:"phone" form:"phone"`
}
