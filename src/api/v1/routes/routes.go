package routes

import (
	"github.com/gofiber/fiber/v3"
	"src/api/v1/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", controllers.Health)

	//v1API := app.Group("/api/v1")
}
