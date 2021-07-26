package models

import (
	"github.com/nikola43/ecoapigorm/models/base"
)

type Clinic struct {
	base.CustomGormModel
	Name      string `gorm:"type:varchar(32) not null" json:"name"`
	Address   string `gorm:"type:varchar(32)" json:"address"`
	Available uint   `gorm:"type:INTEGER not null; DEFAULT:1" json:"available"`
	//EmployeeID       uint           `json:"employee_id"`
	CompanyID uint `json:"company_id"`
	DiskQuote uint `gorm:"type:INTEGER not null; DEFAULT:1073741824" json:"disk_quote"`
	//CreditPrice      uint           `gorm:"type:FLOAT not null; DEFAULT:2.9" json:"credit_price"`
	//ExtendCredits    bool              `gorm:"type:INTEGER not null; DEFAULT:0" json:"extent_credits"`
	AvailableCredits uint       `gorm:"type:INTEGER not null; DEFAULT:0" json:"available_credits"`
	UsedCredits      uint       `gorm:"type:INTEGER not null; DEFAULT:0" json:"used_credits"`
	Clients          []*Client   `gorm:"many2many:clinic_clients;" json:"clinic_clients"`
	Promos           []*Promo    `gorm:"many2many:clinic_promos;" json:"promos"`
	Employees        []Employee `json:"employees"`
}
