package kicks

import (
	"gorm.io/gorm"
	"time"
)

type Kick struct {
	gorm.Model
	Date time.Time   `json:"date"`
	ClientId uint   `json:"client_id"`
}
