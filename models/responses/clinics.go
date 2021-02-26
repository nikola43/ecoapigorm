package models

import (
	"gorm.io/gorm"
)

type Clinic struct {
	gorm.Model
	Name             string   `gorm:"type:varchar(32) not null" json:"name"`
	Address          string   `gorm:"type:varchar(32)" json:"address"`
	Available        uint     `gorm:"type:INTEGER not null" json:"available"`
	EmployeeID       uint     `json:"employee_id"`
	DiskQuote        uint     `gorm:"type:INTEGER not null; DEFAULT:1073741824" json:"disk_quote"`
	AvailableClients uint     `gorm:"type:INTEGER not null; DEFAULT:10" json:"available_users"`
	Clients          []Client `json:"clients"`
	Promos           []Promo  `json:"promos"`
}
