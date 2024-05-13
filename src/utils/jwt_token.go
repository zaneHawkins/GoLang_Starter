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
	token := &T.Tokens{}
	token.AccessToken = createAccessToken(user.ID)
	token.RefreshToken = createRefreshToken(user.ID)
	fmt.Println("Tokens", token)
	return token
}

func createAccessToken(userId string) string {
	var signingKey = []byte(os.Getenv("ACCESS_KEY_SECRET"))

	// Define the claims
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userId,
		"exp":        time.Now().Add(time.Minute * 20).Unix(), // Tokens will expire in 20 min
	}

	return createSignedToken(claims, signingKey)
}

func createRefreshToken(userId string) string {
	var signingKey = []byte(os.Getenv("REFRESH_KEY_SECRET"))

	// Define the claims
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Tokens will expire in 24 hours
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
