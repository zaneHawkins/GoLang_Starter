package controllers

import (
	"github.com/gofiber/fiber/v2"
	"src/config"
)

func Health(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"ok":  1,
		"v":   config.Conf.Version,
		"env": config.Conf.Environment,
	})
}

func SecuredHealth(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"ok":  1,
		"v":   config.Conf.Version,
		"env": config.Conf.Environment,
	})
}
