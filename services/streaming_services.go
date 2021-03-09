package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models/streaming"
)

func GetStreamingByCodeService(code string) (streaming.Streaming, error) {
	var streaming = streaming.Streaming{}

	if err := database.GormDB.Where("code = ?", code).
		First(&streaming).Error;

	err != nil {
		return streaming, err
	}

	return streaming, nil
}
