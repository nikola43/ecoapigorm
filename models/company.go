package models

type Company struct {
	CustomGormModel
	EmployeeID uint       `json:"employee_id"`
	Name       string     `gorm:"type:varchar(32)" json:"name"`
	Employees  []Employee `json:"employees"`
}
