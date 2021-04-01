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

	apiTokenString, err := utils.GeneratePasswordRecoveryToken("client", client.ID)
	if err != nil {
		return err
	}

	userRecovery := recovery.UserRecovery{
		UserId: client.ID,
		Token:  apiTokenString,
		Type:   "client",
	}
	result := database.GormDB.Create(&userRecovery)
	if result.Error != nil {
		return result.Error
	}

	sendEmailManager := utils.SendEmailManager{
		ToEmail:               client.Email,
		ToName:                client.Name,
		RecoveryPasswordToken: apiTokenString,
	}

	sendEmailManager.SendMail("recovery_password.html", "Recuperar contraseña")

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

	apiTokenString, err := utils.GeneratePasswordRecoveryToken("employee", employee.ID)
	if err != nil {
		return err
	}

	userRecovery := recovery.UserRecovery{
		UserId: employee.ID,
		Token:  apiTokenString,
		Type:   "employee",
	}
	result := database.GormDB.Create(&userRecovery)
	if result.Error != nil {
		return result.Error
	}
	sendEmailManager := utils.SendEmailManager{
		ToEmail:               employee.Email,
		ToName:                employee.Name,
		RecoveryPasswordToken: apiTokenString,
	}

	sendEmailManager.SendMail("recovery_password.html", "Recuperar contraseña")

	return nil
}
