package services

import (
	"database/sql"
	Database "github.com/nikola43/ecoapigorm/database"
	models "github.com/nikola43/ecoapigorm/models/responses"
	utils "github.com/nikola43/ecoapigorm/utils"
)

func LoginEmployer(email string, password string) (*models.Employee, error) {
	employer := &models.Employee{}
	dbResult := Database.DB.
		Where("email = ?", email).
		Find(&employer)

	if dbResult.Error != nil {
		return nil, dbResult.Error
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

	employer.Password = utils.HashAndSalt([]byte(employer.Password))
	result := Database.DB.Create(employer)

	if result.Error != nil {
		return nil, result.Error
	}

	employer.Password = ""

	return employer, result.Error
}