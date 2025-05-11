package identity

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterOAuth(app *fiber.App) {
	// OAuth 2.0 endpoints
	app.Get("/authorize", AuthorizationMiddleware)
	app.Post("/token", TokenMiddleware)
}
