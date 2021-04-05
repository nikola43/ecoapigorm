package services

import (
	"errors"
	"fmt"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/payments"
)

func CreatePayment(createPaymentRequest *payments.CreatePaymentRequest) (*models.Payment, error) {
	payment := new(models.Payment)
	clinic := models.Clinic{}
	fmt.Println(createPaymentRequest)
	if err := database.GormDB.First(&clinic, createPaymentRequest.ClinicID).Error; err != nil {
		return nil, err
	}

	if clinic.ID < 1 {
		return nil, errors.New("clinic not found")
	}

	sessionID := ""
	fmt.Println(sessionID)
	for ok := true; ok; ok = len(payment.SessionID) > 0 {
		sessionID = GenerateRandomCode(8)
		database.GormDB.Where("session_id = ?", sessionID).Find(&payment)
	}
	fmt.Println(sessionID)

	payment.CompanyID = createPaymentRequest.CompanyID
	payment.SessionID = sessionID
	payment.EmployeeID = createPaymentRequest.EmployeeID
	payment.ClinicID = createPaymentRequest.ClinicID
	payment.ClinicName = clinic.Name
	payment.Amount = createPaymentRequest.Amount
	payment.Quantity = createPaymentRequest.Quantity
	payment.IsPaid = false

	result := database.GormDB.Create(payment)
	if result.Error != nil {
		return nil, result.Error
	}

	return payment, nil
}

func GetPaymentBySessionID(sessionID string) (*models.Payment, error) {
	payment := new(models.Payment)
	clinic := new(models.Clinic)

	result := database.GormDB.Where("session_id = ?", sessionID).Find(&payment)
	if result.Error != nil {
		return nil, result.Error
	}

	result = database.GormDB.Where("id = ?", payment.ClinicID).Find(&clinic)
	if result.Error != nil {
		return nil, result.Error
	}

	if payment.ID < 1 {
		return nil, errors.New("payment not found")
	}

	return payment, nil
}

func ValidatePayment(sessionID string) (*models.Payment, error) {
	payment := new(models.Payment)
	clinic := new(models.Clinic)

	result := database.GormDB.Where("session_id = ?", sessionID).Find(&payment)
	if result.Error != nil {
		return nil, result.Error
	}

	result = database.GormDB.Where("id = ?", payment.ClinicID).Find(&clinic)
	if result.Error != nil {
		return nil, result.Error
	}

	if payment.ID < 1 {
		return nil, errors.New("payment not found")
	}

	if payment.IsPaid == false {
		payment.IsPaid = true
		database.GormDB.Model(&payment).Update("is_paid", true)
		database.GormDB.Model(&clinic).Update("available_credits", clinic.AvailableCredits+payment.Quantity)
	}

	return payment, nil
}
