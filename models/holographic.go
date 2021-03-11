package models

import "github.com/nikola43/ecoapigorm/models/base"

type Holographic struct {
	base.CustomGormModel
	ClientID     uint   `gorm:"type:INTEGER not null" json:"client_id"`
	Url          string `gorm:"type:varchar(256) not null" json:"url"`
	ThumbnailUrl string `gorm:"type:varchar(256) not null" json:"thumbnail_url"`
	Size         uint   `gorm:"type:INTEGER not null" json:"size"`
}
