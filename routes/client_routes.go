package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/utils"
	jwtware "github.com/gofiber/jwt/v2"

)

func ClientRoutes (router fiber.Router) {
	router.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))


	// routes
	router.Get("/", controllers.LoginClient)
	router.Get("/", controllers.LoginClient)
	router.Get("/", controllers.LoginClient)
	router.Get("/", controllers.LoginClient)
	router.Get("/", controllers.LoginClient)
	//clientRouter.Get("/:id", controller.GetSingleProduct)
}
