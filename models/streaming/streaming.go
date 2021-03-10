package streaming

import "github.com/nikola43/ecoapigorm/models/base"

type Streaming struct {
	base.CustomGormModel
	ClientID uint   `gorm:"type:INTEGER not null" json:"client_id"`
	Code     string `gorm:"type:varchar(5) not null" json:"code"`
	Url      string `gorm:"type:varchar(256) not null" json:"url"`
	Code      string `gorm:"type:varchar(256) not null" json:"code"`
}
