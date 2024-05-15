package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	H "src/handler"
)

func ValidateAccessToken() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")
		if tokenString == "" {
			return H.BuildError(ctx, "Token is required", fiber.StatusUnauthorized, nil)
		}

		// Extract token from Authorization header (Bearer token)
		authToken := tokenString[len("Bearer "):]

		// Parse and validate the token
		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			var signingKey = []byte(os.Getenv("ACCESS_KEY_SECRET"))
			return signingKey, nil
		})
		if err != nil || !token.Valid {
			return H.BuildError(ctx, "Invalid token", fiber.StatusUnauthorized, nil)
		}

		// Token is valid, proceed with the next handler
		return ctx.Next()
	}
}
