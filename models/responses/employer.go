package models

import (
	"gorm.io/gorm"
)

type Employer struct {
	gorm.Model
	Email      string         `json:"email"`
	Password      string         `json:"password"`
	Name          string `json:"name"`
	LastName      string `json:"lastname"`
	Clinics []*Clinic `gorm:"many2many:employer_clinics;"`
}