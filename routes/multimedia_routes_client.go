package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/utils"
)

func MultimediaClientRoutes (router fiber.Router) {
	multimediaRouter := router.Group("/multimedia")

	multimediaClientRouter := multimediaRouter.Group("/client")

	// use middleware
	multimediaClientRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

	// /api/v1/multimedia/client/:client_id/download
	multimediaClientRouter.Get("/:client_id/download", controllers.DownloadAllMultimediaContentByClientID)

	// /api/v1/multimedia/client/:client_id/holographics
	multimediaClientRouter.Get("/:client_id/holographics", controllers.GetAllHolographicsByClientID)

	// /api/v1/multimedia/client/:client_id/images
	multimediaClientRouter.Get("/:client_id/images", controllers.GetAllImagesByClientID)

	// /api/v1/multimedia/client/:client_id/videos
	multimediaClientRouter.Get("/:client_id/videos", controllers.GetAllVideosByClientID)

	// /api/v1/multimedia/client/:client_id/heartbeat
	multimediaClientRouter.Get("/:client_id/heartbeat", controllers.GetHeartbeatByClientID)

	// /api/v1/multimedia/client/:client_id/heartbeat
	multimediaClientRouter.Get("/:client_id/streamings", controllers.GetAllStreamingByClientID)


}
