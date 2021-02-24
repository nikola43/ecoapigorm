package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigormgorm/controllers"
)

func ClientRoutes (router fiber.Router) {

	// routes
	router.Get("/", controllers.GetAllProducts)
	//clientRouter.Get("/:id", controller.GetSingleProduct)
}