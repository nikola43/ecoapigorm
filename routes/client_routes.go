package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/ecoapigorm/controllers"
)

func ClientRoutes (router fiber.Router) {
	// /api/v1/client
	clientRouter := router.Group("/client")


	// /api/v1/client/:client_id/upload
	clientRouter.Post("/:client_id/upload", controllers.UploadMultimedia)

	// /api/v1/client/recovery
	clientRouter.Post("/recovery", controllers.PassRecoveryClient)

	// /api/v1/client/create
	clientRouter.Post("/create", controllers.CreateClient)

	// use jwt
	//clientRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))



	// /api/v1/client/change_password
	clientRouter.Post("/change_password", controllers.ChangePassClient)

	// /api/v1/client/images
	imagesRouter := clientRouter.Group("/images")

	// /api/v1/client/images/:client_id
	imagesRouter.Get("/:client_id", controllers.GetAllImagesByClientID)

	// /api/v1/client/videos
	videosRouter := clientRouter.Group("/videos")

	// /api/v1/client/videos/:client_id
	videosRouter.Get("/:client_id", controllers.GetAllVideosByClientID)

	// /api/v1/client/streaming
	streamingsRouter := clientRouter.Group("/streaming")

	// /api/v1/client/streaming/:client_id
	streamingsRouter.Get("/:client_id", controllers.GetAllVideosByClientID)
}
