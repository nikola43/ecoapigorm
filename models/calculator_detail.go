package models

import "gorm.io/gorm"

type CalculatorDetail struct {
	gorm.Model
	Image              string                   `gorm:"type:varchar(254) not null" json:"image"`
	Text                 string                 `gorm:"type:varchar(254) not null" json:"text"`
}

