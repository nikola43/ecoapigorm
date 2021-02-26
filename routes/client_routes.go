package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/utils"
)

func ClientRoutes (router fiber.Router) {
	// /api/v1/client
	clientRouter := router.Group("/client")
	clientRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

	// /api/v1/client/images
	imagesRouter := clientRouter.Group("/images")

	// /api/v1/client/images
	imagesRouter.Get("/", controllers.GetAllImagesByClientID)
}
