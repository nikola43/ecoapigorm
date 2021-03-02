package services

import (
	"database/sql"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	utils "github.com/nikola43/ecoapigorm/utils"
)

func LoginEmployer(email string, password string) (*models.Employee, error) {
	employer := &models.Employee{}

	GormDBResult := database.GormDB.
		Where("email = ?", email).
		Find(&employer)

	if GormDBResult.Error != nil {
		return nil, GormDBResult.Error
	}

	match := utils.ComparePasswords(employer.Password, []byte(password))
	if match == false {
		return nil, sql.ErrNoRows
	}

	// remove password
	employer.Password = ""

	return employer, nil
}

func CreateNewEmployer(employer *models.Employee) (*models.Employee, error) {
	//TODO validate

	employer.Password = utils.HashPassword([]byte(employer.Password))
	result := database.GormDB.Create(employer)

	if result.Error != nil {
		return nil, result.Error
	}

	employer.Password = ""

	return employer, result.Error
}
