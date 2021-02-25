package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ClinicId             uint                 `json:"clinic_id"`
	Email                string               `gorm:"type:varchar(64) not null" json:"email"`
	Password             string               `gorm:"type:varchar(32) not null" json:"password"`
	Name                 string               `gorm:"type:varchar(32) not null" json:"name"`
	LastName             string               `gorm:"type:varchar(32)" json:"lastname"`
	Videos               []Video              `json:"videos"`
	Images               []Image              `json:"images"`
	Heartbeat            []Heartbeat          `json:"heartbeat"`
	Streaming            Streaming            `json:"streaming"`
	Recovery             Recovery             `json:"recovery"`
	PushNotificationDatas []PushNotificationData `json:"push_notification_datas"`
}
