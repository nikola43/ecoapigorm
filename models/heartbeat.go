package models

type Heartbeat struct {
	CustomGormModel
	ClientID uint   `gorm:"type:INTEGER not null" json:"client_id"`
	Url      string `gorm:"type:varchar(256) not null" json:"url"`
	Size     uint   `gorm:"type:INTEGER not null" json:"size"`
}
