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
	companyRouter := router.Group("/company")


	// use jwt
	companyRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_EMPLOYEE_KEY"))}))

	// check Employee.Role == 'admin'
	companyRouter.Use(middleware.AdminEmployeeMiddleware)

	// /api/v1/company/:company_id/employees
	companyRouter.Get("/:company_id/employees", controllers.GetEmployeesByCompanyID)

	// /api/v1/company/create
	companyRouter.Post("/", controllers.CreateCompany)

	// /api/v1/company/:company_id
	companyRouter.Get("/:company_id", controllers.GetCompanyById)

	// /api/v1/company/:company_id/clinics
	companyRouter.Get("/:company_id/clinics", controllers.GetClinicsByCompanyID)
}
