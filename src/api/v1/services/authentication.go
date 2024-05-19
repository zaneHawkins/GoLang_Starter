package services

import (
	"fmt"
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

func SignUp(db boil.ContextExecutor, ctx *fiber.Ctx, user *M.User) (*T.Tokens, *T.ServiceError) {
	// Encrypt Password
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, H.ServiceError(fiber.StatusInternalServerError, "Password Encryption failed", err)
	}
	user.Password = string(pass)

	// Insert user
	err = user.Insert(ctx.UserContext(), db, boil.Infer())
	if err != nil {
		return nil, H.ServiceError(fiber.StatusConflict, "Email already has an account", err)
	}

	// Get New JWT tokens
	tokens := U.NewTokens(*user)
	return tokens, nil
}

func Login(db boil.ContextExecutor, ctx *fiber.Ctx, loginRequest *T.LoginRequest) (*T.Tokens, *T.ServiceError) {
	// Find User by email
	user, err := M.Users(qm.Where("email = ?", loginRequest.Email)).One(ctx.UserContext(), db)
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

func ForgotPassword(db boil.ContextExecutor, ctx *fiber.Ctx, email string) *T.ServiceError {
	// Find User by email
	user, err := M.Users(qm.Where("email = ?", email)).One(ctx.UserContext(), db)
	if err != nil {
		return H.ServiceError(fiber.StatusNotFound, "User not found", err)
	}

	// Generate Password Reset Token
	resetToken := U.ResetPasswordToken(*user)

	// Send Reset Email
	er := sendPasswordResetEmail(*user, resetToken)
	if er != nil {
		return er
	}

	return nil
}

func RefreshToken(ctx *fiber.Ctx) (*T.Tokens, *T.ServiceError) {
	refreshToken := ctx.Locals("refreshToken").(string)
	userId := ctx.Locals("userId").(string)
	newTokens := U.RefreshToken(userId)
	newTokens.RefreshToken = refreshToken
	return newTokens, nil
}

func sendPasswordResetEmail(user M.User, resetToken string) *T.ServiceError {
	resetURL := fmt.Sprintf("https://yourdomain.com/reset-password?token=%s", resetToken)
	emailBody := fmt.Sprintf("Click the link to reset your password: %s", resetURL)

	email := U.Mail{
		To:      []string{user.Email},
		Subject: "Reset your password",
		Body:    emailBody,
	}

	err := U.SendEmail(&email)
	if err != nil {
		return H.ServiceError(fiber.StatusInternalServerError, "Failed to send email", err)
	}

	return nil
}
