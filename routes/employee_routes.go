package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/middleware"
	"github.com/nikola43/ecoapigorm/utils"
)

func EmployeeRoutes(router fiber.Router) {
	// /api/v1/employee
	employeeRouter := router.Group("/employee")

	// /api/v1/employee/create
	employeeRouter.Post("/", controllers.CreateEmployee)

	// /api/v1/employee/recovery
	employeeRouter.Post("/recovery", controllers.PassRecoveryEmployee)

	// /api/v1/employee/validate_recovery
	employeeRouter.Get("/validate_recovery/:recovery_token", controllers.ValidateRecovery)

	// /api/v1/client/change_password
	employeeRouter.Post("/change_password", controllers.ChangePassEmployee)

	// /api/v1/employee/validate_invitation/:invitation_token
	employeeRouter.Get("/validate_invitation/:invitation_token", controllers.ValidateInvitation)

	// use jwt
	employeeRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

	// check Employee.Role == 'admin'
	employeeRouter.Use(middleware.AdminEmployeeMiddleware)

	// /api/v1/employee/:parent_employee_id/employees
	employeeRouter.Get("/:parent_employee_id/employees", controllers.GetEmployeesByParentEmployeeID)

	// /api/v1/employee/invite
	employeeRouter.Post("/invite", controllers.Invite)

	// /api/v1/employee/:employee_id/companies
	employeeRouter.Get("/:employee_id/companies", controllers.GetCompaniesByEmployeeID) //TODO revisar este servicio, no tiene sentido una lista de compañias

	// /api/v1/employee/:employee_id
	employeeRouter.Delete("/:employee_id", controllers.DeleteEmployeeByEmployeeID)
}
