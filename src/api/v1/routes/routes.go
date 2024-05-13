package routes

import (
	"github.com/gofiber/fiber/v2"
	"src/api/v1/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", controllers.Health)

	APIv1 := app.Group("/api/v1")

	SetAuthenticationRoutes(APIv1)
}
