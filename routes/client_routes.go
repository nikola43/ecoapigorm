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

	// /api/v1/client/:client_id | DELETE
	// clientRouter.Delete("/:client_id", controllers.CreateClient)

	// /api/v1/client/recovery
	clientRouter.Post("/recovery", controllers.PassRecoveryClient)

	// /api/v1/client/validate_recovery
	clientRouter.Get("/validate_recovery/:recovery_token", controllers.ValidateRecovery)

	// use jwt
	clientRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

	// todo implementar cambiar contraseña cliente
	// /api/v1/client/change_password
	clientRouter.Post("/change_password", controllers.ChangePassClient)

	// /api/v1/clinic/refresh
	clientRouter.Get("/:client_id/refresh", controllers.RefreshClient)
}
