package recovery

import "github.com/nikola43/ecoapigorm/models/base"

type UserRecovery struct {
	base.CustomGormModel
	UserId  uint   `gorm:"type:INTEGER not null" json:"user_id"`
	Token string `gorm:"type:varchar(254) not null" json:"token"`
	Type string `gorm:"type:varchar(64) not null" json:"type"`
}
