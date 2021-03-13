package models

import "github.com/nikola43/ecoapigorm/models/base"

type Invitation struct {
	base.CustomGormModel
	EmployeeID uint   `gorm:"type:INTEGER not null" json:"employee_id"`
	Token      string `gorm:"type:varchar(256) not null" json:"token"`
}
