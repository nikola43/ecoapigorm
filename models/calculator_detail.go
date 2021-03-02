package models

import "gorm.io/gorm"

type CalculatorDetail struct {
	gorm.Model
	Week              uint                   `gorm:"type:INTEGER not null" json:"week"`
	Image              string                   `gorm:"type:varchar(254) not null" json:"image"`
	Text                 string                 `gorm:"type:varchar(254) not null" json:"text"`
}
