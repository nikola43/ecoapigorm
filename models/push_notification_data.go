package models

import "github.com/nikola43/ecoapigorm/models/base"

type PushNotificationData struct {
	base.CustomGormModel
	ClientID   uint   `gorm:"type:INTEGER not null" json:"client_id"`
	DeviceType string `gorm:"type:varchar(8) not null" json:"device_type"`
	PushToken  string `gorm:"type:varchar(128) not null" json:"push_token"`
}
