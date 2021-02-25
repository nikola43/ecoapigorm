package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Email    string `gorm:"type:varchar(64)" json:"email"`
	Password string `gorm:"type:varchar(32)" json:"password"`
	Name     string `gorm:"type:varchar(32)" json:"name"`
	LastName string `gorm:"type:varchar(32)" json:"lastname"`
	Clinic   Clinic `json:"clinic"`
	// Clinics []*Clinic `gorm:"many2many:employer_clinics;"`
}
