package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/utils"
)

func ClinicRoutes (router fiber.Router) {
	router.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

	clinics := router.Group("/clinics")

	clinics.Get("/", controllers.GetAllClinics)
}
