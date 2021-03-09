package models

import "github.com/nikola43/ecoapigorm/models/base"

type Image struct {
	base.CustomGormModel
	ClientID uint   `gorm:"type:INTEGER not null" json:"client_id"`
	Url      string `gorm:"type:varchar(256) not null" json:"url"`
	Size     uint   `gorm:"type:INTEGER not null" json:"size"`
}
