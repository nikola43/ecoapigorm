package kicks

import "time"

type AddKickRequest struct {
	ID        uint `json:"id"`
	ClientId        uint `json:"client_id"`
	Date      time.Time `json:"date"`
}
