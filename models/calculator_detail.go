package models

import (
	"github.com/nikola43/ecoapigorm/models/base"
)

type CalculatorDetail struct {
	base.CustomGormModel
	Week  uint   `gorm:"type:INTEGER not null" json:"week"`
	Image string `gorm:"type:varchar(254) not null" json:"image"`
	Title string `gorm:"type:varchar(64) not null" json:"title"`
	Text  string `gorm:"type:varchar(2048) not null" json:"text"`
}
