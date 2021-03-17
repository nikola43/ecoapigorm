package models

import (
	"github.com/nikola43/ecoapigorm/models/base"
)

type Calculator struct {
	base.CustomGormModel
	ClientID uint `gorm:"type:INTEGER not null" json:"client_id"`
	Week     uint `gorm:"type:INTEGER not null" json:"week"`
}
