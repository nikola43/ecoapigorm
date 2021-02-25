package models

import (
	"gorm.io/gorm"
)

type Heartbeat struct {
	gorm.Model
	ClientId uint   `json:"client_id"`
	Url      string `gorm:"type:varchar(128)" json:"url"`
	Size     uint   `json:"size"`
}
