package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/controllers"
)

func ClientRoutes (router fiber.Router) {

	// routes
	router.Get("/", controllers.LoginClient)
	router.Get("/", controllers.LoginClient)
	router.Get("/", controllers.LoginClient)
	router.Get("/", controllers.LoginClient)

	router.Get("/", controllers.LoginClient)
	//clientRouter.Get("/:id", controller.GetSingleProduct)
}
