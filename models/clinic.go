package models

import (
	"github.com/nikola43/ecoapigorm/models/base"
	"github.com/nikola43/ecoapigorm/models/promos"
)

type Clinic struct {
	base.CustomGormModel
	Name             string         `gorm:"type:varchar(32) not null" json:"name"`
	Address          string         `gorm:"type:varchar(32)" json:"address"`
	Available        uint           `gorm:"type:INTEGER not null; DEFAULT:1" json:"available"`
	EmployeeID       uint           `json:"employee_id"`
	DiskQuote        uint           `gorm:"type:INTEGER not null; DEFAULT:1073741824" json:"disk_quote"`
	ExtendCredits    bool           `gorm:"type:INTEGER not null; DEFAULT:0" json:"extent_credits"`
	AvailableCredits uint           `gorm:"type:INTEGER not null; DEFAULT:0" json:"available_credits"`
	Clients          []Client       `json:"clients"`
	Promos           []promos.Promo `json:"promos"`
	Employees         []Employee     `json:"employees"`
}
