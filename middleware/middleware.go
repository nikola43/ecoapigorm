package middleware

import "github.com/gofiber/fiber/v2"

var JwtMiddleware = func(context *fiber.Ctx) error {
	// Set some security headers:
	context.Set("X-XSS-Protection", "1; mode=block")
	context.Set("X-Content-Type-Options", "nosniff")
	context.Set("X-Download-Options", "noopen")
	context.Set("Strict-Transport-Security", "max-age=5184000")
	context.Set("X-Frame-Options", "SAMEORIGIN")
	context.Set("X-DNS-Prefetch-Control", "off")

	// Go to next middleware:
	return context.Next()
}
