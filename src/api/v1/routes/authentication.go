package routes

import (
	"github.com/gofiber/fiber/v2"

	"src/api/v1/controllers"
	mw "src/api/v1/middleware"
	C "src/constants"
)

func SetAuthenticationRoutes(router fiber.Router) {

	router.Post("/login", mw.RateLimit(C.Tier2, 0), controllers.Login)

}
