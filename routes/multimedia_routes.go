package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/utils"
)

func MultimediaRoutes (router fiber.Router) {
	multimediaRouter := router.Group("/multimedia")

	multimediaRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

	// /api/v1/multimedia/images/:id
	multimediaRouter.Delete("/images/:id", controllers.DeleteImage)

	// /api/v1/multimedia/videos/:id
	multimediaRouter.Delete("/videos/:id", controllers.DeleteVideo)

	// /api/v1/multimedia/holographic/:id
	multimediaRouter.Delete("/holographic/:id", controllers.DeleteHolographic)

	// /api/v1/multimedia/heartbeat/:id
	multimediaRouter.Delete("/heartbeat/:id", controllers.DeleteHeartbeat)
}
