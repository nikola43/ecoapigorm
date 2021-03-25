package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/utils"
)

func PromoRoutes(router fiber.Router) {
	promoRouter := router.Group("/promo")

	promoRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

	// /api/v1/promo/create
	promoRouter.Post("/", controllers.CreatePromo)

	// /api/v1/promo/client
	promoRouter.Get("/client", controllers.GetPromosController)

	// /api/v1/promo/:promo_id
	promoRouter.Delete("/:promo_id", controllers.DeletePromoByID)
}
