package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	M "src/models"
	T "src/types"
	"time"
)

func NewTokens(user M.User) *T.Tokens {
	tokens := &T.Tokens{}
	tokens.AccessToken = createAccessToken(user.ID)
	tokens.RefreshToken = createRefreshToken(user.ID)
	return tokens
}

func RefreshToken(userId string) *T.Tokens {
	tokens := &T.Tokens{}
	tokens.AccessToken = createAccessToken(userId)
	return tokens
}

func ResetPasswordToken(user M.User) string {
	return createResetPasswordToken(user.ID, user.Email)
}

func createAccessToken(userId string) string {
	var signingKey = []byte(os.Getenv("ACCESS_KEY_SECRET"))

	// Define the claims
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userId,
		"exp":        time.Now().Add(time.Minute * 20).Unix(), // Token will expire in 20 min
	}

	return createSignedToken(claims, signingKey)
}

func createRefreshToken(userId string) string {
	var signingKey = []byte(os.Getenv("REFRESH_KEY_SECRET"))

	// Define the claims
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token will expire in 24 hours
	}

	return createSignedToken(claims, signingKey)
}

func createResetPasswordToken(userId string, email string) string {
	var signingKey = []byte(os.Getenv("RESET_KEY_SECRET"))

	// Define the claims
	claims := jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // Token will expire in 1 hour
	}

	return createSignedToken(claims, signingKey)
}

func createSignedToken(claims jwt.MapClaims, signingKey []byte) string {
	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return ""
	}

	return tokenString
}
