package models

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	ClientId     uint   `json:"client_id"`
	Url          string `gorm:"type:varchar(128)" json:"url"`
	ThumbnailUrl string `gorm:"type:varchar(128)" json:"thumbnail_url"`
	Size         uint   `json:"size"`
}
