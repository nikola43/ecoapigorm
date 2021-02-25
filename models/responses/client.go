package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ClinicId  uint        `json:"clinic_id"`
	Email     string      `gorm:"type:varchar(64)" json:"email"`
	Password  string      `gorm:"type:varchar(32)" json:"password"`
	Name      string      `gorm:"type:varchar(32)" json:"name"`
	LastName  string      `gorm:"type:varchar(32)" json:"lastname"`
	Videos    []Video     `json:"videos"`
	Images    []Image     `json:"images"`
	Heartbeat []Heartbeat `json:"heartbeat"`
	Streaming Streaming   `json:"streaming"`
}
