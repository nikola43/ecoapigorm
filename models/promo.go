package models

import (
	"github.com/nikola43/ecoapigorm/models/base"
)

type Promo struct {
	base.CustomGormModel
	Title    string   `gorm:"type:varchar(64) not null" json:"title"`
	Text     string   `gorm:"type:varchar(256) not null" json:"text"`
	ImageUrl string   `gorm:"type:varchar(256) not null" json:"image_url"`
	Week     uint     `json:"week"`
	StartAt  string   `gorm:"type:varchar(128) not null" json:"start_at"`
	EndAt    string   `gorm:"type:varchar(128) not null" json:"end_at"`
	Clinics  []Clinic `gorm:"many2many:clinic_promos;" json:"clinics"`
}