package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/recovery"
	"github.com/nikola43/ecoapigorm/utils"
)

func PassRecoveryClientService(request *recovery.PassRecoveryRequest) error {
	client := &models.Client{}

	GormDBResult := database.GormDB.
		Where("email = ?", request.Email).
		Find(&client)

	if GormDBResult.Error != nil {
		return GormDBResult.Error
	}

	apiTokenString, err := utils.GenerateClientToken(client.Email, client.ClinicID, client.ID)
	if err != nil {
		return err
	}

	recovery := recovery.UserRecovery{
		UserId:          client.ID,
		Token:           apiTokenString,
		Type:            "client",
	}
	result := database.GormDB.Create(&recovery)
	if result.Error != nil {
		return result.Error
	}
	SendMailRecovery(client.Email, recovery.Token)

	return nil
}

func PassRecoveryEmployeeService(request *recovery.PassRecoveryRequest) error {
	employee := &models.Employee{}

	GormDBResult := database.GormDB.
		Where("email = ?", request.Email).
		Find(&employee)

	if GormDBResult.Error != nil {
		return GormDBResult.Error
	}

	apiTokenString, err := utils.GenerateClientToken(employee.Email, employee.ClinicID, employee.ID)
	if err != nil {
		return err
	}

	recovery := recovery.UserRecovery{
		UserId:          employee.ID,
		Token:           apiTokenString,
		Type:            "employee",
	}
	result := database.GormDB.Create(&recovery)
	if result.Error != nil {
		return result.Error
	}
	SendMailRecovery(employee.Email, recovery.Token)

	return nil
}
