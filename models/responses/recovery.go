package models

import (
	"gorm.io/gorm"
)

type Recovery struct {
	gorm.Model
	ClientID uint   `gorm:"type:INTEGER not null" json:"client_id"`
	Url      string `gorm:"type:varchar(128) not null" json:"url"`
	Token    string `gorm:"type:varchar(128) not null" json:"token"`
}
