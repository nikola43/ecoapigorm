package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/controllers"
)

func SignUpRoutes (router fiber.Router) {
	authRouter := router.Group("/signup")
	authRouter.Get("/client", controllers.SignUpClient)
	authRouter.Get("/employee", controllers.SignUpEmployee)
}
