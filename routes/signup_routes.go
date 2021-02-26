package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/controllers"
)

func SignUpRoutes(router fiber.Router) {
	// /api/v1/signup
	authRouter := router.Group("/signup")

	// /api/v1/signup/client
	authRouter.Get("/client", controllers.SignUpClient)

	// /api/v1/signup/employee
	authRouter.Get("/employee", controllers.SignUpEmployee)
}
