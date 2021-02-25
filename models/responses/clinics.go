package models

import (
	"gorm.io/gorm"
)

type Clinic struct {
	gorm.Model
	Name       string `gorm:"type:varchar(32)" json:"name"`
	Address    string `gorm:"type:varchar(32)" json:"address"`
	Available  uint   `json:"available"`
	EmployeeId uint   `json:"employee_id"`
	Clients  []Client `json:"clients"`
}
