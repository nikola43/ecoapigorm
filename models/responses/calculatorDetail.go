package models

import "gorm.io/gorm"

type CalculatorDetail struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	ImageUrl string `json:"image_url"`
	Text     string `json:"text"`
}
