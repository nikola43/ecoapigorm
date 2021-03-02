package models

type Recovery struct {
	CustomGormModel
	ClientID uint   `gorm:"type:INTEGER not null" json:"client_id"`
	Token    string `gorm:"type:varchar(128) not null" json:"token"`
}
