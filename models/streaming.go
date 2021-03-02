package models

import (
	"gorm.io/gorm"
)

type Streaming struct {
	gorm.Model
	ClientID uint   `gorm:"type:INTEGER not null" json:"client_id"`
	Url      string `gorm:"type:varchar(256) not null" json:"url"`
}
