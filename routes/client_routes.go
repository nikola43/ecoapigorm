package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/controllers"
)

func ClientRoutes (router fiber.Router) {
	// /api/v1/client
	clientRouter := router.Group("/client")


	// /api/v1/client/:client_id/upload
	clientRouter.Post("/:client_id/upload/:upload_mode", controllers.UploadMultimedia)

	// /api/v1/client/recovery
	clientRouter.Post("/recovery", controllers.PassRecoveryClient)

	// /api/v1/client/create
	clientRouter.Post("/create", controllers.CreateClient)

	// /api/v1/client/:client_id
	clientRouter.Get("/:client_id", controllers.GetClientById)

	// use jwt
	//clientRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))



	// /api/v1/client/change_password
	clientRouter.Post("/change_password", controllers.ChangePassClient)

	// /api/v1/client/:client_id/images
	clientRouter.Get("/:client_id/images", controllers.GetAllImagesByClientID)

	// /api/v1/client/:client_id/videos
	clientRouter.Get("/:client_id/videos", controllers.GetAllVideosByClientID)

	// /api/v1/client/:client_id/holographics
	clientRouter.Get("/:client_id/holographics", controllers.GetAllHolographicsByClientID)

	// /api/v1/client/:client_id/heartbeat
	clientRouter.Get("/:client_id/heartbeat", controllers.GetHeartbeatByClientID)

	// /api/v1/client/:client_id/heartbeat
	clientRouter.Get("/:client_id/streamings", controllers.GetAllStreamingByClientID)

	// /api/v1/client/:client_id/download
	clientRouter.Get("/:client_id/download", controllers.DownloadAllMultimediaContentByClientID)
}
