package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/middleware"
	"github.com/nikola43/ecoapigorm/utils"
)

func ClinicRoutes (router fiber.Router) {
	// /api/v1/clinic
	clinicRouter := router.Group("/clinic")

	// use jwt
	clinicRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

	// /api/v1/clinic/:clinic_id/clients
	clinicRouter.Get("/:clinic_id/clients", controllers.GetClientsByClinicID)

	// check Employee.Role == 'admin'
	clinicRouter.Use(middleware.AdminEmployeeMiddleware)

	// /api/v1/clinic/create
	clinicRouter.Post("/create", controllers.CreateClinic)

	// /api/v1/clinic/create
	clinicRouter.Get("/:clinic_id", controllers.GetClinicById)
}
