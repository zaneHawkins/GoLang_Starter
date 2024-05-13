package services

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
	M "src/models"
	T "src/types"
	U "src/utils"
)

func Login(db boil.ContextExecutor, ctx context.Context, body *T.LoginRequest) (*T.Tokens, *T.ServiceError) {
	// Find User by email
	user, err := M.FindUser(ctx, db, body.Email)
	if err != nil {
		return nil, &T.ServiceError{
			Message: "User not found",
			Error:   err,
			Code:    fiber.StatusNotFound,
		}
	}

	// Check Password matches
	if !comparePasswords(user.Password, []byte(body.Password)) {
		return nil, &T.ServiceError{
			Message: "Incorrect Password",
			Code:    fiber.StatusForbidden,
		}
	}

	// Get New JWT tokens
	tokens := U.NewTokens(*user)

	return tokens, nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
