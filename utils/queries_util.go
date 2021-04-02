package utils

import (
	"fmt"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/promos"
	"github.com/nikola43/ecoapigorm/models/recovery"
	streamingModels "github.com/nikola43/ecoapigorm/models/streamings"
)


func GetModelByField(dest interface{}, fieldName string, fieldValue interface{}) error {
	var model interface{}

	// todo crear todos los casos
	switch dest.(type) {
	case *models.Client:
		model = dest.(*models.Client)
	case *models.Clinic:
		model = dest.(*models.Clinic)
	case *models.Employee:
		model = dest.(*models.Employee)
	case *streamingModels.Streaming:
		model = dest.(*streamingModels.Streaming)
	case *promos.Promo:
		model = dest.(*promos.Promo)
	case *models.Image:
		model = dest.(*models.Image)
	case *models.Video:
		model = dest.(*models.Video)
	case *models.Holographic:
		model = dest.(*models.Holographic)
	case *models.Heartbeat:
		model = dest.(*models.Heartbeat)
	case *models.Invitation:
		model = dest.(*models.Invitation)
	case *recovery.UserRecovery:
		model = dest.(*recovery.UserRecovery)
	}

	result := database.GormDB.Where(fieldName+" = ?", fieldValue).First(model)
	if result != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil
}
