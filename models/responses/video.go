package models

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	ClientID     uint   `gorm:"type:INTEGER not null" json:"client_id"`
	Url          string `gorm:"type:varchar(128) not null" json:"url"`
	ThumbnailUrl string `gorm:"type:varchar(128) not null" json:"thumbnail_url"`
	Size         uint   `gorm:"type:INTEGER not null" json:"size"`
}
