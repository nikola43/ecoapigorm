package kicks

import (
	"github.com/nikola43/ecoapigorm/models/base"
	"time"
)

type Kick struct {
	base.CustomGormModel
	Date time.Time   `json:"date"`
	ClientId uint   `json:"client_id"`
}
