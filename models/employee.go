package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	EmployeeID    uint          `json:"employee_id"`
	Email         string        `gorm:"index; unique; type:varchar(64) not null" json:"email"`
	Password      string        `gorm:"type:varchar(256) not null" json:"password"`
	Name          string        `gorm:"type:varchar(32) not null" json:"name"`
	Phone         string        `json:"phone" xml:"phone" form:"phone" validate:"required"`
	LastName      string        `gorm:"type:varchar(32)" json:"lastname"`
	Role          string        `gorm:"type:varchar(32) not null; DEFAULT:'employee'" json:"role"`
	Clinic        Clinic        `json:"clinic"`
	PaymentMethod PaymentMethod `json:"payment_method"`
}
