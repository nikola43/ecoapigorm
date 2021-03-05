package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/middleware"
	"github.com/nikola43/ecoapigorm/utils"
)

func CompanyRoutes (router fiber.Router) {
	// /api/v1/company
	clinicRouter := router.Group("/company")

	// use jwt
	clinicRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

	// /api/v1/company/:company_id/employees
	clinicRouter.Get("/:company_id/employees", controllers.GetEmployeesByCompanyID)

	// check Employee.Role == 'admin'
	clinicRouter.Use(middleware.AdminEmployeeMiddleware)

	// /api/v1/company/create
	clinicRouter.Post("/create", controllers.CreateCompany)

	// /api/v1/company/:company_id
	clinicRouter.Get("/:company_id", controllers.GetCompanyById)
}
