package models


type employerClinic struct {
	EmployerID uint64 `gorm:"primaryKey;autoIncrement:false"`
	ClinicID     uint64 `gorm:"primaryKey;autoIncrement:false"`
}
