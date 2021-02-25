package models

import (
	"gorm.io/gorm"
)

type Promo struct {
	gorm.Model
	ClinicId uint   `gorm:"type:INTEGER not null" json:"clinic_id"`
	Title    string `gorm:"type:varchar(128) not null" json:"title"`
	Text     string `gorm:"type:varchar(128) not null" json:"text"`
	ImageUrl string `gorm:"type:varchar(128) not null" json:"image_url"`
	StartAt  string `gorm:"type:varchar(128) not null" json:"start_at"`
	EndAt    string `gorm:"type:varchar(128) not null" json:"end_at"`
}
