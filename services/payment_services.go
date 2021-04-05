package services

import (
	"errors"
	"fmt"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/payments"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func CreatePayment(createPaymentRequest *payments.CreatePaymentRequest) (*models.Payment, error) {
	stripe.Key = "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"

	payment := new(models.Payment)
	clinic := models.Clinic{}
	fmt.Println(createPaymentRequest)
	if err := database.GormDB.First(&clinic, createPaymentRequest.ClinicID).Error; err != nil {
		return nil, err
	}

	if clinic.ID < 1 {
		return nil, errors.New("clinic not found")
	}

	/*
		fmt.Println(sessionID)
		for ok := true; ok; ok = len(payment.SessionID) > 0 {
			sessionID = GenerateRandomCode(8)
			database.GormDB.Where("session_id = ?", sessionID).Find(&payment)
		}
		fmt.Println(sessionID)
	*/

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("eur"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Cliente"),
					},
					UnitAmount: stripe.Int64(3),
				},
				Quantity: stripe.Int64(int64(createPaymentRequest.Quantity)),
			},
		},
		SuccessURL: stripe.String("https://panel.mimatrona.stelast.es/success-payment"),
		CancelURL:  stripe.String("https://panel.mimatrona.stelast.es/error-payment"),
	}

	checkoutSession, err := session.New(params)
	if err != nil {
		return nil, err
	}

	payment.CompanyID = createPaymentRequest.CompanyID
	payment.SessionID = checkoutSession.ID
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

	c, _ := session.Get(
		payment.SessionID,
		nil,
	)
	fmt.Println("c")
	fmt.Println(c.PaymentStatus)

	if payment.IsPaid == false && c.PaymentStatus == stripe.CheckoutSessionPaymentStatusPaid {
		payment.IsPaid = true
		database.GormDB.Model(&payment).Update("is_paid", true)
		database.GormDB.Model(&clinic).Update("available_credits", clinic.AvailableCredits+payment.Quantity)
	}

	return payment, nil
}
