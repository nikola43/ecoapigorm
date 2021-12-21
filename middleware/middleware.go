package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/utils"
)

func WebSocketUpgradeMiddleware(context *fiber.Ctx) error {
	// IsWebSocketUpgrade returns true if the client
	// requested upgrade to the WebSocket protocol.
	if websocket.IsWebSocketUpgrade(context) {
		context.Locals("allowed", true)
		return context.Next()
	}

	return fiber.ErrUpgradeRequired
}

func XApiKeyMiddleware(context *fiber.Ctx) error {
	requestApiKey := context.Get("XAPIKEY")
	serverApiKey := utils.GetEnvVariable("XAPIKEY")
	fmt.Println("requestApiKey")
	fmt.Println(requestApiKey)
	fmt.Println("serverApiKey")
	fmt.Println(serverApiKey)
	// context.h

	if requestApiKey != serverApiKey {
		return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "unauthorized",
		})
	}

	return context.Next()
}

func AdminEmployeeMiddleware(context *fiber.Ctx) error {
	employeeTokenClaims, err := utils.GetEmployeeTokenClaims(context)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusUnauthorized, err, context)
	}

	if employeeTokenClaims.Role != models.ADMIN_ROLE {
		return utils.ReturnErrorResponse(fiber.StatusUnauthorized, errors.New("employee not is admin"), context)
	}

	return context.Next()
}

func EmployeeEmployeeMiddleware(context *fiber.Ctx) error {
	employeeTokenClaims, err := utils.GetEmployeeTokenClaims(context)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusUnauthorized, err, context)
	}

	if employeeTokenClaims.Role != models.EMPLOYEE_ROLE {
		return utils.ReturnErrorResponse(fiber.StatusUnauthorized, errors.New("user not is employee"), context)
	}

	return context.Next()
}
