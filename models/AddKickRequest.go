package models

import "time"

type AddKickRequest struct {
	ID        uint `json:"id"`
	Date      time.Time `json:"date"`
}
