package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/controllers"
)

func AuthRoutes (router fiber.Router) {
	authRouter := router.Group("/auth")
	authRouter.Post("/client", controllers.LoginClient)
	authRouter.Post("/employee", controllers.LoginEmployee)
}
