package models

import "gorm.io/gorm"

type ConfigClinic struct {
	gorm.Model
	ClinicID uint
}

