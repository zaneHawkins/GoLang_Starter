package services

import (
	"context"
	"github.com/friendsofgo/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
	H "src/handler"
	M "src/models"
	T "src/types"
	U "src/utils"
)

func SignUp(db boil.ContextExecutor, ctx context.Context, user *M.User) (*T.Tokens, *T.ServiceError) {
	// Encrypt Password
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, H.ServiceError(fiber.StatusInternalServerError, "Password Encryption failed", err)
	}
	user.Password = string(pass)

	// Insert user
	err = user.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, H.ServiceError(fiber.StatusConflict, "Email already has an account", err)
	}

	// Get New JWT tokens
	tokens := U.NewTokens(*user)
	return tokens, nil
}

func Login(db boil.ContextExecutor, ctx context.Context, loginRequest *T.LoginRequest) (*T.Tokens, *T.ServiceError) {
	// Find User by email
	user, err := M.Users(qm.Where("email = ?", loginRequest.Email)).One(ctx, db)
	if err != nil {
		return nil, H.ServiceError(fiber.StatusNotFound, "User not found", err)
	}

	// Check Password matches
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return nil, H.ServiceError(fiber.StatusForbidden, "Password did not match", errors.New("Incorrect Password"))
	}

	// Get New JWT tokens
	tokens := U.NewTokens(*user)
	return tokens, nil
}
