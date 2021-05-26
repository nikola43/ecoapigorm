package controllers

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/clients"
	clinicModels "github.com/nikola43/ecoapigorm/models/clinic"
	"github.com/nikola43/ecoapigorm/models/promos"
	streamingModels "github.com/nikola43/ecoapigorm/models/streamings"
	"github.com/nikola43/ecoapigorm/services"
	"github.com/nikola43/ecoapigorm/utils"
	"strconv"
)

func GetCreditsClinicById(context *fiber.Ctx) error {
	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)
	var credits uint = 0
	var err error

	if credits, err = services.GetCreditsClinicById(uint(clinicID)); err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "clinic not found",
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"available_credits": credits,
	})
}

func GetClinicById(context *fiber.Ctx) error {
	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)
	clinic := &models.Clinic{}
	var err error

	if clinic, err = services.GetClinicByID(uint(clinicID)); err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "clinic not found",
		})
	}

	return context.JSON(clinic)
}

func CreateClinic(context *fiber.Ctx) error {
	createClinicRequest := new(clinicModels.CreateClinicRequest)
	createClinicResponse := new(clinicModels.CreateClinicResponse)
	var err error
	clinic := models.Clinic{}

	employeeTokenClaims, err := utils.GetEmployeeTokenClaims(context)
	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	// parse request

	err = context.BodyParser(createClinicRequest)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	// validation ---------------------------------------------------------------------
	v := validator.New()
	err = v.Struct(createClinicRequest)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e != nil {
				return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"error": "validation_error: " + e.Field(),
				})
			}
		}
	}

	// check if clinic already exist
	result := database.GormDB.
		Where("name = ?", createClinicRequest.Name).
		Find(&clinic)

	if result.Error == nil {

	}
	if clinic.ID > 0 {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": errors.New("clinic already exist"),
		})
	}

	// create and response
	if createClinicResponse, err = services.CreateClinic(employeeTokenClaims.CompanyID, createClinicRequest);
		err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.JSON(createClinicResponse)

}

func GetClientsByClinicID(context *fiber.Ctx) error {
	clientsList := make([]clients.ListClientResponse, 0)
	var err error

	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)
	fmt.Println(clinicID)
	if clientsList, err = services.GetClientsByClinicID(uint(clinicID))
		err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "clinic not found",
		})
	} else {
		return context.JSON(clientsList)
	}
}

func CreateClientFromClinic(context *fiber.Ctx) error {
	createClientRequest := new(clients.CreateClientRequest)
	listClientResponse := new(clients.ListClientResponse)
	var err error

	// parse request
	if err = context.BodyParser(createClientRequest);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	// validation ---------------------------------------------------------------------
	v := validator.New()
	err = v.Struct(createClientRequest)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e != nil {
				return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"error": "validation_error: " + e.Field(),
				})
			}
		}
	}

	// check if client already exist
	client := models.Client{}
	GormDBResult := database.GormDB.
		Where("email = ?", createClientRequest.Email).
		Find(&client)

	if GormDBResult.Error != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": "internal server",
		})
	}

	// create and response
	if listClientResponse, err = services.CreateClientFromClinic(createClientRequest);
		err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	} else {
		return context.JSON(listClientResponse)
	}
}

func GetAllStreamingByClinicID(context *fiber.Ctx) error {
	id := context.Params("clinic_id")
	streamings := make([]streamingModels.Streaming, 0)
	var err error

	if streamings, err = services.GetAllStreamingByClinicID(id); err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(streamings)
}

func GetAllPromosByClinicID(context *fiber.Ctx) error {
	id := context.Params("clinic_id")
	promos := make([]promos.Promo, 0)
	var err error

	if promos, err = services.GetAllPromosByClinicID(id); err != nil {
		return context.SendStatus(fiber.StatusInternalServerError)
	}

	return context.Status(fiber.StatusOK).JSON(promos)
}

func GetEmployeesByClinicID(context *fiber.Ctx) error {
	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)

	if list, err := services.GetEmployeesByClinicID(uint(clinicID)); err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "Company not found",
		})
	} else {
		return context.JSON(list)
	}
}

func UpdateClinicByID(context *fiber.Ctx) error {
	clinic := new(models.Clinic)

	// parse request
	err := context.BodyParser(clinic)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	// todo validate
	// validation ---------------------------------------------------------------------

	list, err := services.UpdateClinic(clinic)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return context.JSON(list)

}


func UpdateCredits(context *fiber.Ctx) error {
	clinic := new(models.Clinic)

	// parse request
	err := context.BodyParser(clinic)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	updatedClinic, err := services.UpdateClinic(clinic)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return context.JSON(updatedClinic)

}


func LinkClient(context *fiber.Ctx) error {
	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)

	err := services.LinkClient(uint(clientID), uint(clinicID))
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func DeleteClinicByID(context *fiber.Ctx) error {
	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)

	err := services.DeleteClinicByID(uint(clinicID))
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func GetPromosByWeekAndClinicID(context *fiber.Ctx) error {
	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)
	week, _ := strconv.ParseUint(context.Params("week"), 10, 64)

	promosList, err := services.GetPromosByWeekAndClinicID(uint(week), uint(clinicID))
	if err != nil {
		return context.SendStatus(fiber.StatusBadRequest)
	}

	return context.Status(fiber.StatusOK).JSON(promosList)
}

func GetClientById(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)
	clinicID, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)

	client, err := services.GetClientById(uint(clinicID), uint(clientID))
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(client)

}
