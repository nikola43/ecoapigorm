package models

import (
	"gorm.io/gorm"
	"time"
)

type ClinicClient struct {
	ClinicID  uint `gorm:"primaryKey"`
	ClientID uint `gorm:"primaryKey"`
	DiskQuote        uint           `gorm:"type:INTEGER not null; DEFAULT:1073741824" json:"disk_quote"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
