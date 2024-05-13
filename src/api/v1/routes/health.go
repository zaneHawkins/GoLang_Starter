package routes

import (
	"github.com/gofiber/fiber/v2"
	"src/api/v1/controllers"
	mw "src/api/v1/middleware"
	C "src/constants"
)

func SetHealthRoutes(router fiber.Router) {

	router.Get("/health", mw.RateLimit(C.Tier5, 0), controllers.Health)
	router.Get("/secured-health", mw.RateLimit(C.Tier5, 0), mw.ValidateToken(), controllers.SecuredHealth)

}
