package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/utils"
)

func EmployeeRoutes (router fiber.Router) {
	// /api/v1/employee
	employeeRouter := router.Group("/employee")

	// /api/v1/employee/create
	employeeRouter.Post("/create", controllers.CreateEmployee)

	employeeRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

}
