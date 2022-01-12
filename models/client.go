package models

import (
	"github.com/nikola43/ecoapigorm/models/base"
	streamingModels "github.com/nikola43/ecoapigorm/models/streamings"
)

type Client struct {
	base.CustomGormModel
	//ClinicID              uint                   `gorm:"type:INTEGER NULL; DEFAULT:NULL" json:"clinic_id" xml:"clinic_id" form:"clinic_id"`
	Email                 string                      `gorm:"index; unique; type:varchar(64) not null" json:"email"`
	Password              string                      `gorm:"type:varchar(256) not null; size:256" json:"password" xml:"password" form:"password"`
	Name                  string                      `gorm:"type:varchar(32) not null" json:"name" xml:"name" form:"name"`
	Phone                 string                      `gorm:"type:varchar(32) not null" json:"phone" xml:"phone" form:"phone"`
	LastName              string                      `gorm:"type:varchar(32)" json:"last_name" xml:"last_name" form:"last_name"`
	PregnancyDate         base.CustomNullTime         `json:"pregnancy_date" xml:"pregnancy_date"`
	Videos                []Video                     `json:"videos" xml:"videos" form:"videos"`
	Holographs            []Holographic               `json:"holographics" xml:"holographics" form:"holographics"`
	Images                []Image                     `json:"images" xml:"images" form:"images"`
	Heartbeat             []Heartbeat                 `json:"heartbeat" xml:"heartbeat" form:"heartbeat"`
	Streaming             []streamingModels.Streaming `json:"streamings" xml:"streamings" form:"streamings"`
	Clinics               []*Clinic                   `gorm:"many2many:clinic_clients;" json:"clinics"`
	PushNotificationDatas []PushNotificationData      `json:"push_notification_datas" xml:"push_notification_datas" form:"push_notification_datas"`
	ProfileColor          string                      `gorm:"type:varchar(16) not null" json:"profile_color" xml:"profile_color" form:"profile_color"`
}
