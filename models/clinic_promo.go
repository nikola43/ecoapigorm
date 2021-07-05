package models

import (
	"gorm.io/gorm"
	"time"
)

type ClinicPromo struct {
	ClinicID  uint           `gorm:"primaryKey" json:"clinic_id"`
	PromoID   uint           `gorm:"primaryKey" json:"promo_id"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
