package utils

import (
	"fmt"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
)

func GetModelByField(dest interface{}, fieldName string, fieldValue interface{}) interface{} {
	var model interface{}

	// todo crear todos los casos
	switch dest.(type) {
	case *models.Client:
		model = dest.(*models.Client)
	case *models.Clinic:
		model = dest.(*models.Clinic)
	}

	err := database.GormDB.Where(fieldName+" = ?", fieldValue).First(model)
	if err != nil {
		fmt.Println(err.Error)
		return nil
	}

	return model
}
