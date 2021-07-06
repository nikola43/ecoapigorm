package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/utils"
)

func MultimediaClinicRoutes (router fiber.Router) {
	multimediaRouter := router.Group("/multimedia")

	multimediaClinicRouter := multimediaRouter.Group("/clinic")

	// use middleware
	multimediaClinicRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_EMPLOYEE_KEY"))}))


	// /api/v1/multimedia/images/:id
	multimediaRouter.Delete("/images/:id", controllers.DeleteImage)

	// /api/v1/multimedia/videos/:id
	multimediaRouter.Delete("/videos/:id", controllers.DeleteVideo)

	// /api/v1/multimedia/holographic/:id
	multimediaRouter.Delete("/holographic/:id", controllers.DeleteHolographic)

	// /api/v1/multimedia/heartbeat/:id
	multimediaRouter.Delete("/heartbeat/:id", controllers.DeleteHeartbeat)

	// /api/v1/multimedia/clinic/:clinic_id/client/:client_id/upload/:upload_mode
	multimediaClinicRouter.Post("/:clinic_id/client/:client_id/upload/:upload_mode", controllers.UploadMultimedia)

	// /api/v1/multimedia/clinic/:clinic_id/promo/:promo_id/upload
	multimediaClinicRouter.Post("/:clinic_id/promo/:promo_id/upload", controllers.UploadPromoImage)

	// /api/v1/multimedia/clinic/:clinic_id/:client_id/images
	multimediaClinicRouter.Get("/:clinic_id/client/:client_id/images", controllers.GetAllImagesByClientAndClinicID)

	multimediaClinicRouter.Get("/:clinic_id/client/:client_id/videos", controllers.GetAllVideosByClientAndClinicID)

	multimediaClinicRouter.Get("/:clinic_id/client/:client_id/heartbeat", controllers.GetHeartbeatByClientAndClinicID)

	multimediaClinicRouter.Get("/:clinic_id/client/:client_id/streamings", controllers.GetAllStreamingByClientANDClinicID)

	multimediaClinicRouter.Get("/:clinic_id/client/:client_id/holographics", controllers.GetAllHolographicsByClientID)
}
