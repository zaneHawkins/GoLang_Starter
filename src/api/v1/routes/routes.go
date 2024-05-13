package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	SetHealthRoutes(app)

	APIv1 := app.Group("/api/v1")
	SetAuthenticationRoutes(APIv1)
}
