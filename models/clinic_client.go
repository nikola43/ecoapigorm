package models

import (
	"gorm.io/gorm"
	"time"
)

type ClinicClient struct {
	ClinicID       uint `gorm:"primaryKey"`
	ClientID       uint `gorm:"primaryKey"`
	DiskQuoteLevel uint `gorm:"type:INTEGER not null; DEFAULT:1" json:"disk_quote_level"`
	CreatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
