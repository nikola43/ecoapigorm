package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/controllers"
)

func AuthRoutes (router fiber.Router) {
	authRouter := router.Group("/auth")
	authRouter.Get("/login_client", controllers.LoginClient)
	authRouter.Get("/login_employee", controllers.LoginEmployee)
}