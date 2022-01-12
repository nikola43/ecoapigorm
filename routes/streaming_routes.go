package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	_ "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/utils"
)

func StreamingRoutes(router fiber.Router) {
	streamingRouter := router.Group("/streaming")

	// /api/v1/streaming/:code
	streamingRouter.Get("/:code", controllers.GetStreamingByCodeController)

	streamingRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_EMPLOYEE_KEY"))}))

	// /api/v1/streaming
	streamingRouter.Post("/", controllers.CreateStreaming)

	// /api/v1/streaming
	streamingRouter.Patch("/", controllers.UpdateStreaming)

	// /api/v1/streaming
	streamingRouter.Delete("/:streaming_id", controllers.DeleteStreamingByID)
}
