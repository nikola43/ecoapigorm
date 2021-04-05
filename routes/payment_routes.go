package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/nikola43/ecoapigorm/controllers"
	"github.com/nikola43/ecoapigorm/utils"
)

func PaymentRoutes(router fiber.Router) {
	promoRouter := router.Group("/payment")

	promoRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

	// /api/v1/payment
	promoRouter.Post("/", controllers.CreatePayment)

	// /api/v1/paymentv2
	promoRouter.Post("/", controllers.CreateCheckoutSession)

	// /api/v1/payment/:session_id
	promoRouter.Get("/:session_id", controllers.GetPaymentBySessionID)

	// /api/v1/payment/:session_id
	promoRouter.Get("/:session_id/validate", controllers.ValidatePayment)
}
