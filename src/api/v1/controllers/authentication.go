package controllers

import (
	"github.com/gofiber/fiber/v2"
	S "src/api/v1/services"
	H "src/handler"
	T "src/types"
	U "src/utils"
)

func Login(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	body := &T.LoginRequest{}

	if err := ctx.BodyParser(body); err != nil {
		return H.BuildError(ctx, "Invalid body", fiber.StatusBadRequest, err)
	}
	user, serviceErr := S.Login(dbTrx, ctx.UserContext(), body)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"ok":       1,
		"products": user,
	})
}
