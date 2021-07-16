package models

import (
	"gorm.io/gorm"
	"time"
)

type ClinicPromo struct {
	ClinicID  uint `gorm:"primaryKey"`
	PromoID   uint `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
