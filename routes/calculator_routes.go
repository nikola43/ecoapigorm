package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/controllers"
)

func CalculatorRoutes (router fiber.Router) {
	// /api/v1/calculator
	calculatorRouter := router.Group("/calculator")

	// /api/v1/calculator/:week
	calculatorRouter.Post("/:week", controllers.LoginClient)
}