package models

import (
	"github.com/nikola43/ecoapigorm/models/base"
)

type Invitation struct {
	base.CustomGormModel
	ParentEmployeeID uint   `json:"parent_employee_id"`
	FromClinicID     uint   `json:"clinic_id"`
	FromEmail        string `gorm:"type:varchar(256) not null" json:"from_email"`
	ToEmail          string `gorm:"type:varchar(256) not null" json:"to_email"`
	Token            string `gorm:"type:varchar(256) not null" json:"token"`
}
