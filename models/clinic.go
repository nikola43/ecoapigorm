package models

type Clinic struct {
	CustomGormModel
	Name             string   `gorm:"type:varchar(32) not null" json:"name"`
	Address          string   `gorm:"type:varchar(32)" json:"address"`
	Available        uint     `gorm:"type:INTEGER not null; DEFAULT:1" json:"available"`
	EmployeeID       uint     `json:"employee_id"`
	DiskQuote        uint     `gorm:"type:INTEGER not null; DEFAULT:1073741824" json:"disk_quote"`
	ExtendClients    bool     `gorm:"type:INTEGER not null; DEFAULT:0" json:"extent_clients"`
	AvailableClients uint     `gorm:"type:INTEGER not null; DEFAULT:10" json:"available_clients"`
	Clients          []Client `json:"clients"`
	Promos           []Promo  `json:"promos"`
}
