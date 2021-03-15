package routes

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
)

func StreamingRoutes(router fiber.Router) {
	streamingRouter := router.Group("/streaming")

	// /api/v1/streaming/:code
	streamingRouter.Get("/:code", controllers.GetStreamingByCodeController)

	// /api/v1/streaming/create
	streamingRouter.Post("/", controllers.CreateStreaming)

	// /api/v1/streaming
	streamingRouter.Delete("/:streaming_id", controllers.DeleteStreamingByID)
}
