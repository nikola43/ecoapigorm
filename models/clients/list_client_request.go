package clients

type ListClientRequest struct {
	ID         uint   `json:"id"  xml:"id" form:"id"`
	ClinicID   uint   `json:"clinic_id"  xml:"clinic_id" form:"clinic_id"`
	ClinicName string `json:"clinic_name"  xml:"clinic_name" form:"clinic_name"`
	Email      string `json:"email" xml:"email" form:"email"`
	Name       string `json:"name"  xml:"name" form:"name"`
	LastName   string `json:"last_name" xml:"last_name" form:"last_name"`
	Phone      string `json:"phone" xml:"phone" form:"phone"`
	Week       uint   `json:"week" xml:"week" form:"week"`
	CreatedAt  string `json:"created_at" xml:"created_at" form:"created_at"`
}
