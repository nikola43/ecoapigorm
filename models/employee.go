package models

import "github.com/nikola43/ecoapigorm/models/base"

type Employee struct {
	base.CustomGormModel
	ParentEmployeeID uint      `json:"parent_employee_id"`
	Company          Company   `json:"company"`
	CompanyID        uint      `gorm:"type:INTEGER NULL; DEFAULT:NULL" json:"company_id" xml:"company_id" form:"company_id"`
	Email            string    `gorm:"index; unique; type:varchar(64) not null" json:"email"`
	Password         string    `gorm:"type:varchar(256) not null" json:"password"`
	Name             string    `gorm:"type:varchar(32) not null" json:"name"`
	Phone            string    `json:"phone" xml:"phone" form:"phone" validate:"required"`
	LastName         string    `gorm:"type:varchar(32)" json:"lastname"`
	IsFirstLogin     bool      `json:"is_first_login"`
	Role             string    `gorm:"type:varchar(32) not null; DEFAULT:'employee'" json:"role"`
	Clinic           Clinic    `json:"clinic"`
	Payment          []Payment `json:"payment" xml:"payment" form:"payment"`
}
