package routes

import (
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/utils"
	"github.com/gofiber/fiber/v2"
)

func UploadRoutes(uploadRouter fiber.Router) {

	uploadRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_EMPLOYEE_KEY"))}))

	// /api/v1/multimedia/clinic/:clinic_id/client/:client_id/upload/:upload_mode
	uploadRouter.Post("/multimedia/:clinic_id/client/:client_id/upload/:upload_mode", controllers.UploadMultimedia)
}
