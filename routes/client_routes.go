package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/utils"
)

func ClientRoutes(router fiber.Router) {
	// /api/v1/client
	clientRouter := router.Group("/client")

	// /api/v1/client | CREATE
	clientRouter.Post("/", controllers.CreateClient)

	// /api/v1/client/:client_id | READ
	clientRouter.Get("/:client_id", controllers.GetClientById)

	// /api/v1/client/:client_email | READ
	clientRouter.Get("/:client_email/exist", controllers.GetClientByEmail)

	// /api/v1/client/:client_id | DELETE
	// clientRouter.Delete("/:client_id", controllers.CreateClient)

	// /api/v1/client/recovery
	clientRouter.Post("/recovery", controllers.PassRecoveryClient)

	// use jwt
	clientRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))


	// /api/v1/clinic/refresh
	clientRouter.Get("/:client_id/refresh", controllers.RefreshClient)

	// /api/v1/client/change_password
	clientRouter.Post("/change_password", controllers.ChangePassClient)

	// /api/v1/client/:client_id
	clientRouter.Delete("/:client_id", controllers.UnassignClientByID)

	// /api/v1/client/:client_id | UPDATE
	clientRouter.Patch("/:client_id", controllers.UpdateClient)
}
