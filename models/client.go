package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ClinicID              uint                   `gorm:"type:INTEGER null; DEFAULT:NULL" json:"clinic_id" xml:"clinic_id" form:"clinic_id"`
	Email                 string                 `gorm:"type:varchar(64) not null" json:"email" xml:"email" form:"email"`
	Password              string                 `gorm:"type:varchar(256) not null; size:256" json:"password" xml:"password" form:"password"`
	Name                  string                 `gorm:"type:varchar(32) not null" json:"name" xml:"name" form:"name"`
	Phone                 string                 `gorm:"type:varchar(32) not null" json:"phone" xml:"phone" form:"phone"`
	LastName              string                 `gorm:"type:varchar(32)" json:"lastname" xml:"lastname" form:"lastname"`
	Videos                []Video                `json:"videos" xml:"videos" form:"videos"`
	Images                []Image                `json:"images" xml:"images" form:"images"`
	Heartbeat             []Heartbeat            `json:"heartbeat" xml:"heartbeat" form:"heartbeat"`
	Streaming             Streaming              `json:"streaming" xml:"streaming" form:"streaming"`
	Recovery              Recovery               `json:"recovery" xml:"recovery" form:"recovery"`
	PushNotificationDatas []PushNotificationData `json:"push_notification_datas" xml:"push_notification_datas" form:"push_notification_datas"`
}
