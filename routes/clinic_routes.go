package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/controllers"
)

func ClinicRoutes (router fiber.Router) {
	clinics := router.Group("/clinics")

	clinics.Get("/", controllers.GetAllClinics)
}