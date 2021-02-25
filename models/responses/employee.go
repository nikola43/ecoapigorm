package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Email     string `gorm:"type:varchar(64) not null" json:"email"`
	Password  string `gorm:"type:varchar(32) not null" json:"password"`
	Name      string `gorm:"type:varchar(32) not null" json:"name"`
	LastName  string `gorm:"type:varchar(32)" json:"lastname"`
	Role      string `gorm:"type:varchar(32) not null; DEFAULT:'employee'" json:"role"`
	Clinic    Clinic `json:"clinic"`
}
