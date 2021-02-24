package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigormgorm/routes"
)

func main() {

	app := fiber.New()
	api := app.Group("/api")      // /api
	v1 := api.Group("/v1")        // /api/v1

	//app.Use(middleware.Logger())

	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// use routes
	routes.ClientRoutes(v1)
	routes.ClinicRoutes(v1)

	app.Listen(":3000")
}
