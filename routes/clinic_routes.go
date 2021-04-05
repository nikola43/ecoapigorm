package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/middleware"
	"github.com/nikola43/ecoapigorm/utils"
)

func ClinicRoutes(router fiber.Router) {
	// /api/v1/clinic
	clinicRouter := router.Group("/clinic")

	// use jwt
	clinicRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

	// /api/v1/clinic/:clinic_id/clients
	clinicRouter.Get("/:clinic_id/clients", controllers.GetClientsByClinicID)

	// /api/v1/clinic/create_client
	clinicRouter.Post("/create_client", controllers.CreateClientFromClinic)

	// /api/v1/clinic/:clinic_id/client/:client_id | READ
	clinicRouter.Get("/:clinic_id/client/:client_id", controllers.GetClientById)

	// /api/v1/clinic/:clinic_id/client/:client_email/exist | READ
	clinicRouter.Get("/:clinic_id/client/:client_email/exist", controllers.GetClientClinicIDByEmail)

	// /api/v1/clinic/:clinic_id/streamings
	clinicRouter.Get("/:clinic_id/streamings", controllers.GetAllStreamingByClinicID)

	// /api/v1/clinic/:clinic_id/promos
	clinicRouter.Get("/:clinic_id/promos", controllers.GetAllPromosByClinicID)

	// /api/v1/clinic/:clinic_id/:session_id/buy_credits
	clinicRouter.Get("/:clinic_id/:session_id/buy_credits", controllers.BuyCredits)

	// /api/v1/clinic/:clinic_id/:session_id/buy_credits
	clinicRouter.Get("/:clinic_id/link/:client_id", controllers.LinkClient)

	// /api/v1/clinic/:clinic_id/
	clinicRouter.Delete("/:clinic_id", controllers.DeleteClinicByID)

	// /api/v1/clinic/:clinic_id/promos/:week
	clinicRouter.Get("/:clinic_id/promos/:week", controllers.GetPromosByWeekAndClinicID)

	// /api/v1/clinic/:clinic_id/employees
	clinicRouter.Get("/:clinic_id/employees", controllers.GetEmployeesByClinicID)


	// /api/v1/clinic/:clinic_id
	clinicRouter.Get("/:clinic_id/credits", controllers.GetCreditsClinicById)

	// check Employee.Role == 'admin'
	clinicRouter.Use(middleware.AdminEmployeeMiddleware)

	// /api/v1/clinic/create
	clinicRouter.Post("/", controllers.CreateClinic)

	// /api/v1/clinic/:clinic_id
	clinicRouter.Get("/:clinic_id", controllers.GetClinicById)

	// /api/v1/clinic
	clinicRouter.Patch("/", controllers.UpdateClinicByID)
}
