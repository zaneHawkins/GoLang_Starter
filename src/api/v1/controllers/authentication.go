package controllers

import (
	"github.com/friendsofgo/errors"
	"github.com/gofiber/fiber/v2"
	S "src/api/v1/services"
	H "src/handler"
	M "src/models"
	T "src/types"
	U "src/utils"
)

func SignUp(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	body := &M.User{}

	if err := ctx.BodyParser(body); err != nil {
		return H.BuildError(ctx, "Invalid body", fiber.StatusBadRequest, err)
	}

	tokens, serviceErr := S.SignUp(dbTrx, ctx, body)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"tokens": tokens,
	})

}

func Login(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	body := &T.LoginRequest{}

	if err := ctx.BodyParser(body); err != nil {
		return H.BuildError(ctx, "Invalid body", fiber.StatusBadRequest, err)
	}

	tokens, serviceErr := S.Login(dbTrx, ctx, body)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"tokens": tokens,
	})
}

func RefreshToken(ctx *fiber.Ctx) error {

	tokens, serviceErr := S.RefreshToken(ctx)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"tokens": tokens,
	})
}

func ForgotPassword(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	email := ctx.Query("email")

	if email == "" {
		return H.BuildError(ctx, "Missing email parameter", fiber.StatusBadRequest, errors.New("Missing email parameter"))
	}

	serviceErr := S.ForgotPassword(dbTrx, ctx, email)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, nil)
}

func ResetPassword(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	body := &T.LoginRequest{}

	if err := ctx.BodyParser(body); err != nil {
		return H.BuildError(ctx, "Invalid body", fiber.StatusBadRequest, err)
	}

	if body.Email != "" && body.Email != ctx.Locals("email").(string) {
		return H.BuildError(ctx, "Invalid request", fiber.StatusBadRequest, errors.New("Token does not match requested email"))
	}

	body.Email = ctx.Locals("email").(string)

	serviceErr := S.ResetPassword(dbTrx, ctx, body)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, nil)
}
