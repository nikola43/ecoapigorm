package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/utils"
)

func MultimediaRoutes (router fiber.Router) {
	multimediaRouter := router.Group("/multimedia")
	multimediaClientRouter := multimediaRouter.Group("/client")
	multimediaClinicRouter := multimediaRouter.Group("/clinic")

	// use middleware
	multimediaRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

	// /api/v1/multimedia/images/:id
	multimediaRouter.Delete("/images/:id", controllers.DeleteImage)

	// /api/v1/multimedia/videos/:id
	multimediaRouter.Delete("/videos/:id", controllers.DeleteVideo)

	// /api/v1/multimedia/holographic/:id
	multimediaRouter.Delete("/holographic/:id", controllers.DeleteHolographic)

	// /api/v1/multimedia/heartbeat/:id
	multimediaRouter.Delete("/heartbeat/:id", controllers.DeleteHeartbeat)

	// CLIENT

	// /api/v1/multimedia/client/:client_id/images
	multimediaClientRouter.Get("/:client_id/images", controllers.GetAllImagesByClientID)

	// /api/v1/multimedia/client/:client_id/videos
	multimediaClientRouter.Get("/:client_id/videos", controllers.GetAllVideosByClientID)

	// /api/v1/multimedia/client/:client_id/heartbeat
	multimediaClientRouter.Get("/:client_id/heartbeat", controllers.GetHeartbeatByClientID)

	// /api/v1/multimedia/client/:client_id/heartbeat
	multimediaClientRouter.Get("/:client_id/streamings", controllers.GetAllStreamingByClientID)

	// /api/v1/multimedia/clinic/:clinic_id/client/:client_id/upload/:upload_mode
	multimediaClinicRouter.Post("/:clinic_id/client/:client_id/upload/:upload_mode", controllers.UploadMultimedia)

	// /api/v1/multimedia/clinic/:clinic_id/:client_id/images
	multimediaClinicRouter.Get("/:clinic_id/client/:client_id/images", controllers.GetAllImagesByClientAndClinicID)

	multimediaClinicRouter.Get("/:clinic_id/client/:client_id/videos", controllers.GetAllVideosByClientAndClinicID)

	multimediaClinicRouter.Get("/:clinic_id/client/:client_id/heartbeat", controllers.GetHeartbeatByClientAndClinicID)

	multimediaClinicRouter.Get("/:clinic_id/client/:client_id/streamings", controllers.GetAllStreamingByClientANDClinicID)

	multimediaClinicRouter.Get("/:clinic_id/client/:client_id/holographics", controllers.GetAllHolographicsByClientID)


	// /api/v1/multimedia/client/:client_id/holographics
	multimediaClientRouter.Get("/:client_id/holographics", controllers.GetAllHolographicsByClientID)

	// /api/v1/multimedia/client/:client_id/download
	multimediaClientRouter.Get("/:client_id/download", controllers.DownloadAllMultimediaContentByClientID)
}
