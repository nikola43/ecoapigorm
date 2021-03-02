package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
)

func GetCalculator(week uint) (*models.CalculatorDetail, error) {
	calculator := &models.CalculatorDetail{}
	res := database.GormDB.
		Find(&calculator,week)

	return calculator, res.Error
}

