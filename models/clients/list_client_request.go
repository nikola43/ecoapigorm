package clients

import (
	"github.com/nikola43/ecoapigorm/models/base"
	"time"
)

type ListClientResponse struct {
	ID             uint                `json:"id"  xml:"id" form:"id"`
	ClinicID       uint                `json:"clinic_id"  xml:"clinic_id" form:"clinic_id"`
	Email          string              `json:"email" xml:"email" form:"email"`
	Name           string              `json:"name"  xml:"name" form:"name"`
	LastName       string              `json:"last_name" xml:"last_name" form:"last_name"`
	Phone          string              `json:"phone" xml:"phone" form:"phone"`
	CreatedAt      time.Time           `json:"created_at" xml:"created_at" form:"created_at"`
	PregnancyDate  base.CustomNullTime `json:"pregnancy_date" xml:"pregnancy_date"`
	UsedSize       uint                `json:"used_size" xml:"used_size" form:"used_size"`
	DiskQuoteLevel uint                `gorm:"type:INTEGER not null; DEFAULT:1" json:"disk_quote_level"`
}
