package controllers

import (
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
	tokens, serviceErr := S.SignUp(dbTrx, ctx.UserContext(), body)

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
	tokens, serviceErr := S.Login(dbTrx, ctx.UserContext(), body)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"tokens": tokens,
	})
}
