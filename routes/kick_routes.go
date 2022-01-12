package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/utils"
)

func KickRoutes (router fiber.Router) {
	// /api/v1/client/{client_id}
	kickRouter := router.Group("/client/:client_id")

	kickRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

	// /api/v1/client/{client_id}/kicks
	kickRouter.Get("/kicks", controllers.GetKicksByClientID)

	// /api/v1/client/{client_id}/kicks
	kickRouter.Post("/kicks", controllers.AddKickToClient)

	// /api/v1/client/{client_id}/kicks/{kick_id}
	kickRouter.Delete("/kicks/:kick_id", controllers.DeleteKick)

	// /api/v1/client/{client_id}/kicks/reset
	kickRouter.Post("/kicks/reset", controllers.ResetAllKicksByClientID)
}
