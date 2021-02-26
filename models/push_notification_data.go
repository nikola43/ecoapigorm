package models

import (
	"gorm.io/gorm"
)

type PushNotificationData struct {
	gorm.Model
	ClientID   uint   `gorm:"type:INTEGER not null" json:"client_id"`
	DeviceType string `gorm:"type:varchar(8) not null" json:"device_type"`
	PushToken  string `gorm:"type:varchar(128) not null" json:"push_token"`
}
