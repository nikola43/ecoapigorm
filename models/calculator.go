package models

type Calculator struct {
	CustomGormModel
	ClientID uint `gorm:"type:INTEGER not null" json:"client_id"`
	Week     uint `gorm:"type:INTEGER not null" json:"week"`
}
