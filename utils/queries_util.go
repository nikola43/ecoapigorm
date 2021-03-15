package utils

import (
	"fmt"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
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
	}

	result := database.GormDB.Where(fieldName+" = ?", fieldValue).First(model)
	if result != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil
}
